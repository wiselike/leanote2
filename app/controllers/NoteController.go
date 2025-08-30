package controllers

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/wiselike/revel"
	"gopkg.in/mgo.v2/bson"

	"github.com/wiselike/leanote2/app/info"
	. "github.com/wiselike/leanote2/app/lea"
	"github.com/wiselike/leanote2/app/service"
)

type Note struct {
	BaseController
}

// 笔记首页, 判断是否已登录
// 已登录, 得到用户基本信息(notebook, shareNotebook), 跳转到index.html中
// 否则, 转向登录页面
func (c Note) Index(noteId, online string) revel.Result {
	c.SetLocale()
	userInfo := c.GetUserAndBlogUrl()

	userId := userInfo.UserId.Hex()

	// 没有登录
	if userId == "" {
		return c.Redirect("/login")
	}

	c.ViewArgs["openRegister"] = configService.IsOpenRegister()

	// 已登录了, 那么得到所有信息
	notebooks := notebookService.GetNotebooks(userId)
	shareNotebooks, sharedUserInfos := shareService.GetShareNotebooks(userId)

	// 还需要按时间排序(DESC)得到notes
	notes := []info.Note{}
	noteContent := info.NoteContent{}

	if len(notebooks) > 0 {
		// noteId是否存在
		// 是否传入了正确的noteId
		hasRightNoteId := false
		if IsObjectId(noteId) {
			note := noteService.GetNoteById(noteId)

			if note.NoteId != "" {
				var noteOwner = note.UserId.Hex()
				noteContent = noteService.GetNoteContent(noteId, noteOwner)

				hasRightNoteId = true
				c.ViewArgs["curNoteId"] = noteId
				c.ViewArgs["curNotebookId"] = note.NotebookId.Hex()

				// 打开的是共享的笔记, 那么判断是否是共享给我的默认笔记
				if noteOwner != c.GetUserId() {
					if shareService.HasReadPerm(noteOwner, c.GetUserId(), noteId) {
						// 不要获取notebook下的笔记
						// 在前端下发请求
						c.ViewArgs["curSharedNoteNotebookId"] = note.NotebookId.Hex()
						c.ViewArgs["curSharedUserId"] = noteOwner
						// 没有读写权限
					} else {
						hasRightNoteId = false
					}
				} else {
					_, notes = noteService.ListNotes(c.GetUserId(), note.NotebookId.Hex(), false, c.GetPage(), 50, defaultSortField, false, false)

					// 如果指定了某笔记, 则该笔记放在首位
					lenNotes := len(notes)
					if lenNotes > 1 {
						notes2 := make([]info.Note, len(notes))
						notes2[0] = note
						i := 1
						for _, note := range notes {
							if note.NoteId.Hex() != noteId {
								if i == lenNotes { // 防止越界
									break
								}
								notes2[i] = note
								i++
							}
						}
						notes = notes2
					}
				}
			}

			// 得到最近的笔记
			_, latestNotes := noteService.ListNotes(c.GetUserId(), "", false, c.GetPage(), 50, defaultSortField, false, false)
			c.ViewArgs["latestNotes"] = latestNotes
		}

		// 没有传入笔记
		// 那么得到最新笔记
		if !hasRightNoteId {
			_, notes = noteService.ListNotes(c.GetUserId(), "", false, c.GetPage(), 50, defaultSortField, false, false)
			if len(notes) > 0 {
				noteContent = noteService.GetNoteContent(notes[0].NoteId.Hex(), userId)
				c.ViewArgs["curNoteId"] = notes[0].NoteId.Hex()
			}
		}
	}

	// 当然, 还需要得到第一个notes的content
	//...
	c.ViewArgs["isAdmin"] = configService.GetAdminUsername() == userInfo.Username

	c.ViewArgs["userInfo"] = userInfo
	c.ViewArgs["notebooks"] = notebooks
	c.ViewArgs["shareNotebooks"] = shareNotebooks // note信息在notes列表中
	c.ViewArgs["sharedUserInfos"] = sharedUserInfos

	c.ViewArgs["notes"] = notes
	c.ViewArgs["noteContentJson"] = noteContent
	c.ViewArgs["noteContent"] = noteContent.Content

	c.ViewArgs["tags"] = tagService.GetTags(c.GetUserId())

	c.ViewArgs["globalConfigs"] = configService.GetGlobalConfigForUser()

	// return c.RenderTemplate("note/note.html")

	if isDev, _ := revel.Config.Bool("mode.dev"); isDev && online == "" {
		return c.RenderTemplate("note/note-dev.html")
	} else {
		return c.RenderTemplate("note/note.html")
	}
}

// 首页, 判断是否已登录
// 已登录, 得到用户基本信息(notebook, shareNotebook), 跳转到index.html中
// 否则, 转向登录页面
func (c Note) ListNotes(notebookId string) revel.Result {
	_, notes := noteService.ListNotes(c.GetUserId(), notebookId, false, c.GetPage(), pageSize, defaultSortField, false, false)
	return c.RenderJSON(notes)
}

// 得到trash
func (c Note) ListTrashNotes() revel.Result {
	_, notes := noteService.ListNotes(c.GetUserId(), "", true, c.GetPage(), pageSize, defaultSortField, false, false)
	return c.RenderJSON(notes)
}

// 得到note和内容
func (c Note) GetNoteAndContent(noteId string) revel.Result {
	return c.RenderJSON(noteService.GetNoteAndContent(noteId, c.GetUserId()))
}

func (c Note) GetNoteAndContentBySrc(src string) revel.Result {
	noteId, noteAndContent := noteService.GetNoteAndContentBySrc(src, c.GetUserId())
	ret := info.Re{}
	if noteId != "" {
		ret.Ok = true
		ret.Item = noteAndContent
	}
	return c.RenderJSON(ret)
}

// 得到内容
func (c Note) GetNoteContent(noteId string) revel.Result {
	noteContent := noteService.GetNoteContent(noteId, c.GetUserId())
	return c.RenderJSON(noteContent)
}

// 这里不能用json, 要用post
func (c Note) UpdateNoteOrContent(noteOrContent info.NoteOrContent) revel.Result {
	// 新添加note, 不会创建“历史记录”
	if noteOrContent.IsNew {
		userId := c.GetObjectUserId()
		//		myUserId := userId
		// 为共享新建?
		if noteOrContent.FromUserId != "" {
			userId = bson.ObjectIdHex(noteOrContent.FromUserId)
		}

		note := info.Note{UserId: userId,
			NoteId:     bson.ObjectIdHex(noteOrContent.NoteId),
			NotebookId: bson.ObjectIdHex(noteOrContent.NotebookId),
			Title:      noteOrContent.Title,
			Src:        noteOrContent.Src, // 来源
			Tags:       strings.Split(noteOrContent.Tags, ","),
			Desc:       noteOrContent.Desc,
			ImgSrc:     noteOrContent.ImgSrc,
			IsBlog:     noteOrContent.IsBlog,
			IsMarkdown: noteOrContent.IsMarkdown,
		}
		noteContent := info.NoteContent{NoteId: note.NoteId,
			UserId:       userId,
			IsBlog:       note.IsBlog,
			IsAutoBackup: noteOrContent.IsAutoBackup,
			Content:      noteOrContent.Content,
			Abstract:     noteOrContent.Abstract}

		noteImageService.ReOrganizeImageFiles(c.GetUserId(), noteOrContent.NoteId, noteOrContent.Title, &noteContent.Content, false)
		attachService.ReOrganizeAttachFiles(c.GetUserId(), noteOrContent.NoteId, noteOrContent.Title)

		note = noteService.AddNoteAndContentForController(note, noteContent, c.GetUserId())
		return c.RenderJSON(note)
	}

	noteUpdate := bson.M{}
	needUpdateNote := false

	// Desc前台传来
	if c.Has("Desc") {
		needUpdateNote = true
		noteUpdate["Desc"] = noteOrContent.Desc
	}
	if c.Has("ImgSrc") {
		needUpdateNote = true
		noteUpdate["ImgSrc"] = noteOrContent.ImgSrc
	}
	if c.Has("Title") {
		needUpdateNote = true
		noteUpdate["Title"] = noteOrContent.Title
	}

	if c.Has("Tags") {
		needUpdateNote = true
		noteUpdate["Tags"] = strings.Split(noteOrContent.Tags, ",")
	}

	// web端不控制
	if needUpdateNote {
		noteService.UpdateNote(c.GetUserId(), noteOrContent.NoteId, noteUpdate, -1)
	}

	if c.Has("Title") && !c.Has("Content") {
		noteImageService.ReOrganizeImageFiles(c.GetUserId(), noteOrContent.NoteId, noteOrContent.Title, &noteOrContent.Content, true)
		attachService.ReOrganizeAttachFiles(c.GetUserId(), noteOrContent.NoteId, noteOrContent.Title)
	} else if c.Has("Content") {
		if !c.Has("Title") {
			noteOrContent.Title = noteService.GetNote(noteOrContent.NoteId, c.GetUserId()).Title
		}
		noteImageService.ReOrganizeImageFiles(c.GetUserId(), noteOrContent.NoteId, noteOrContent.Title, &noteOrContent.Content, false)
		attachService.ReOrganizeAttachFiles(c.GetUserId(), noteOrContent.NoteId, noteOrContent.Title)
	}

	//-------------
	// afterContentUsn := 0
	// contentOk := false
	// contentMsg := ""
	if c.Has("Content") {
		// contentOk, contentMsg, afterContentUsn =
		noteService.UpdateNoteContent(c.GetUserId(),
			noteOrContent.NoteId, noteOrContent.Content, noteOrContent.Abstract,
			needUpdateNote, noteOrContent.IsAutoBackup, -1, time.Now())
	} else if !noteOrContent.IsAutoBackup {
		// 更新一下noteContent的状态（更新为手动保存的状态）
		noteService.UpdateAutoBackupState(noteOrContent.NoteId, c.GetUserId(), noteOrContent.IsAutoBackup)
	}

	// Log("usn", "afterContentUsn", afterContentUsn + "")
	// Log(contentOk)
	// Log(contentMsg)
	return c.RenderJSON(true)
}

// 删除note/ 删除别人共享给我的笔记
// userId 是note.UserId
func (c Note) DeleteNote(noteIds []string, isShared bool) revel.Result {
	if !isShared {
		for _, noteId := range noteIds {
			trashService.DeleteNote(noteId, c.GetUserId())
		}
		return c.RenderJSON(true)
	}

	for _, noteId := range noteIds {
		trashService.DeleteSharedNote(noteId, c.GetUserId())
	}

	return c.RenderJSON(true)
}

// 删除trash, 已弃用, 用DeleteNote
func (c Note) DeleteTrash(noteId string) revel.Result {
	return c.RenderJSON(trashService.DeleteTrash(noteId, c.GetUserId()))
}

// 移动note
func (c Note) MoveNote(noteIds []string, notebookId string) revel.Result {
	userId := c.GetUserId()
	for _, noteId := range noteIds {
		noteService.MoveNote(noteId, notebookId, userId)
	}
	return c.RenderJSON(true)
}

// 复制note
func (c Note) CopyNote(noteIds []string, notebookId string) revel.Result {
	copyNotes := make([]info.Note, len(noteIds))
	userId := c.GetUserId()
	for i, noteId := range noteIds {
		copyNotes[i] = noteService.CopyNote(noteId, notebookId, userId)
	}
	re := info.NewRe()
	re.Ok = true
	re.Item = copyNotes
	return c.RenderJSON(re)
}

// 复制别人共享的笔记给我
func (c Note) CopySharedNote(noteIds []string, notebookId, fromUserId string) revel.Result {
	copyNotes := make([]info.Note, len(noteIds))
	userId := c.GetUserId()
	for i, noteId := range noteIds {
		copyNotes[i] = noteService.CopySharedNote(noteId, notebookId, fromUserId, userId)
	}
	re := info.NewRe()
	re.Ok = true
	re.Item = copyNotes
	return c.RenderJSON(re)
}

// ------------
// search
// 通过title搜索
func (c Note) SearchNote(key string) revel.Result {
	_, blogs := noteService.SearchNote(key, c.GetUserId(), c.GetPage(), pageSize, "UpdatedTime", false, false)
	return c.RenderJSON(blogs)
}

// 通过tags搜索
func (c Note) SearchNoteByTags(tags []string) revel.Result {
	_, blogs := noteService.SearchNoteByTags(tags, c.GetUserId(), c.GetPage(), pageSize, "UpdatedTime", false)
	return c.RenderJSON(blogs)
}

// 设置/取消Blog; 置顶
func (c Note) SetNote2Blog(noteIds []string, isBlog, isTop bool) revel.Result {
	for _, noteId := range noteIds {
		noteService.ToBlog(c.GetUserId(), noteId, isBlog, isTop)
	}
	return c.RenderJSON(true)
}
func (c Note) toPdf(note *info.Note) ([]byte, error) {
	noteId := note.NoteId.Hex()
	noteUserId := note.UserId.Hex()
	content := noteService.GetNoteContent(noteId, noteUserId)
	userInfo := userService.GetUserInfo(noteUserId)

	// 将 content 的图片地址替换为 base64
	contentStr := content.Content
	// markdown
	if note.IsMarkdown {
		regImageMarkdown := regexp.MustCompile(`!\[.*?\]\(\s*/(file/outputImage|api/file/getImage)\?fileId=([a-z0-9A-Z]+)\)`)

		findsImageMarkdown := regImageMarkdown.FindAllStringSubmatch(contentStr, -1)
		fmt.Printf("%#v\n", findsImageMarkdown)
		for _, md := range findsImageMarkdown {
			if len(md) < 3 {
				continue
			}
			fileId := md[2]

			// 得到base64编码文件
			fileBase64 := fileService.GetImageBase64(noteUserId, fileId)
			if fileBase64 == "" {
				continue
			}

			// 构造新的 Markdown 图片（去除 alt）
			newMD := "![](" + fileBase64 + ")"

			// 用新内容替换原整段匹配（只替一次，避免误伤）
			contentStr = strings.Replace(contentStr, md[0], newMD, 1)
		}
	} else {
		// 1) 找到每个 <img ...> 标签（不跨标签）
		reImgTag := regexp.MustCompile(`<img\b[^>]*>`)
		// 2) 在标签字符串中查找 fileId（只认相对路径里的 fileId，和 src/data-mce-src 等属性名无关）
		reFileId := regexp.MustCompile(`/(?:file/outputImage|api/file/getImage|api/getImage)\?fileId=([a-z0-9A-Z]+)`)
		// 3) 删除标签内所有 src="..." 或 src='...'（不会误伤 srcset、data-src 等）
		reSrcAttr := regexp.MustCompile(`\s+src\s*=\s*["'][^"']*["']`)
		// 4) 把 <img 开头替换为 <img src="...base64..."，保持其它属性不变
		reImgOpen := regexp.MustCompile(`^<img\b`)
		imgTags := reImgTag.FindAllString(contentStr, -1)
		for _, tag := range imgTags {
			// 找 fileId（从任意属性里的相对路径提取）
			idMatch := reFileId.FindStringSubmatch(tag)
			if len(idMatch) < 2 {
				// 这个 <img> 不符合我们的 fileId 规则，跳过
				continue
			}
			fileId := idMatch[1]

			// 用 fileId 拿到 base64
			fileBase64 := fileService.GetImageBase64(noteUserId, fileId)
			if fileBase64 == "" {
				continue
			}

			// 删除该标签里的所有 src=... 属性
			withoutSrc := reSrcAttr.ReplaceAllString(tag, "")
			// 在 <img 后面插入一个新的 src="base64"
			newTag := reImgOpen.ReplaceAllString(withoutSrc, `<img src="`+fileBase64+`"`)

			// 把新标签回写到正文（一次），保留其它属性（alt/width/height/data-mce-src 等）
			contentStr = strings.Replace(contentStr, tag, newTag, 1)
		}
	}

	if note.Tags != nil && len(note.Tags) > 0 && note.Tags[0] != "" {
	} else {
		note.Tags = nil
	}
	c.ViewArgs["blog"] = note
	c.ViewArgs["content"] = contentStr
	c.ViewArgs["userInfo"] = userInfo
	c.ViewArgs["userBlog"] = blogService.GetUserBlog(noteUserId)
	c.ViewArgs["staticBase"] = configService.GetSiteUrl()

	return c.TemplateOutput("file/pdf.html")
}

// 导出成PDF
func (c Note) ExportPdf(noteId string) revel.Result {
	re := info.NewRe()
	userId := c.GetUserId()
	note := noteService.GetNoteById(noteId)
	if note.NoteId == "" {
		re.Msg = "No Note"
		return c.RenderText("error")
	}

	noteUserId := note.UserId.Hex()
	// 是否有权限
	if noteUserId != userId {
		// 是否是有权限协作的
		if !note.IsBlog && !shareService.HasReadPerm(noteUserId, userId, noteId) {
			re.Msg = "No Perm"
			return c.RenderText("No Perm")
		}
	}

	htmlBytes, err := c.toPdf(&note)
	if err != nil {
		return c.RenderText("template error: " + err.Error())
	}

	// outPdfPath 判断是否需要重新生成
	guid := Md5(string(htmlBytes))
	fileUrlPath := "export_pdf"
	dir := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], fileUrlPath)
	if !MkdirAll(dir) {
		return c.RenderText("error, no dir")
	}
	filename := guid + ".pdf"
	outPdfPath := dir + "/" + filename

	if !IsFileExist(outPdfPath) {
		binPath := configService.GetGlobalStringConfig("exportPdfBinPath")
		// 默认路径
		if binPath == "" {
			if runtime.GOOS == "windows" {
				binPath = `C:\Program Files\wkhtmltopdf\bin\wkhtmltopdf.exe`
			} else {
				binPath = "/usr/local/bin/wkhtmltopdf"
			}
		}

		cmd := exec.Command(binPath, "--lowquality", "--window-status", "done", "--quiet", "-", outPdfPath)
		cmd.Stdin = bytes.NewReader(htmlBytes)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return c.RenderText(fmt.Sprintf("export pdf error(%s): %s", err, output))
		}
	}

	file, err := os.Open(outPdfPath)
	if err != nil {
		return c.RenderText("export pdf error. " + fmt.Sprintf("%v", err))
	}
	// http://stackoverflow.com/questions/8588818/chrome-pdf-display-duplicate-headers-received-from-the-server
	//	filenameReturn = strings.Replace(filenameReturn, ",", "-", -1)
	filenameReturn := note.Title
	filenameReturn = FixFilename(filenameReturn)
	if filenameReturn == "" {
		filenameReturn = "Untitled.pdf"
	} else {
		filenameReturn += ".pdf"
	}
	return c.RenderBinary(file, filenameReturn, revel.Attachment, time.Now()) // revel.Attachment
}
