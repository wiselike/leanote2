package service

import (
	// "fmt"
	"sort"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/wiselike/leanote-of-unofficial/app/db"
	"github.com/wiselike/leanote-of-unofficial/app/info"
	. "github.com/wiselike/leanote-of-unofficial/app/lea"
	//	"html"
)

// 笔记本

type NotebookService struct {
}

// 排序
func sortSubNotebooks(eachNotebooks info.SubNotebooks) info.SubNotebooks {
	// 遍历子, 则子往上进行排序
	for _, eachNotebook := range eachNotebooks {
		if eachNotebook.Subs != nil && len(eachNotebook.Subs) > 0 {
			eachNotebook.Subs = sortSubNotebooks(eachNotebook.Subs)
		}
	}

	// 子排完了, 本层排
	sort.Sort(&eachNotebooks)
	return eachNotebooks
}

// 整理(成有关系)并排序
// GetNotebooks()调用
// ShareService调用
func ParseAndSortNotebooks(userNotebooks []info.Notebook, noParentDelete, needSort bool) info.SubNotebooks {
	// 整理成info.Notebooks
	// 第一遍, 建map
	// notebookId => info.Notebooks
	userNotebooksMap := make(map[bson.ObjectId]*info.Notebooks, len(userNotebooks))
	for _, each := range userNotebooks {
		newNotebooks := info.Notebooks{Subs: info.SubNotebooks{}}
		newNotebooks.NotebookId = each.NotebookId
		newNotebooks.Title = each.Title
		//		newNotebooks.Title = html.EscapeString(each.Title)
		newNotebooks.Title = strings.Replace(strings.Replace(each.Title, "<script>", "", -1), "</script", "", -1)
		newNotebooks.Seq = each.Seq
		newNotebooks.UserId = each.UserId
		newNotebooks.ParentNotebookId = each.ParentNotebookId
		newNotebooks.NumberNotes = each.NumberNotes
		newNotebooks.IsTrash = each.IsTrash
		newNotebooks.IsBlog = each.IsBlog

		// 存地址
		userNotebooksMap[each.NotebookId] = &newNotebooks
	}

	// 第二遍, 追加到父下

	// 需要删除的id
	needDeleteNotebookId := map[bson.ObjectId]bool{}
	for id, each := range userNotebooksMap {
		// 如果有父, 那么追加到父下, 并剪掉当前, 那么最后就只有根的元素
		if each.ParentNotebookId.Hex() != "" {
			if userNotebooksMap[each.ParentNotebookId] != nil {
				userNotebooksMap[each.ParentNotebookId].Subs = append(userNotebooksMap[each.ParentNotebookId].Subs, each) // Subs是存地址
				// 并剪掉
				// bug
				needDeleteNotebookId[id] = true
				// delete(userNotebooksMap, id)
			} else if noParentDelete {
				// 没有父, 且设置了要删除
				needDeleteNotebookId[id] = true
				// delete(userNotebooksMap, id)
			}
		}
	}

	// 第三遍, 得到所有根
	final := make(info.SubNotebooks, len(userNotebooksMap)-len(needDeleteNotebookId))
	i := 0
	for id, each := range userNotebooksMap {
		if !needDeleteNotebookId[id] {
			final[i] = each
			i++
		}
	}

	// 最后排序
	if needSort {
		return sortSubNotebooks(final)
	}
	return final
}

// 得到 notebook 的分类及父分类 ["Life": {}, "Life1": {}, "Life11": {}] 层级关系：Life -> Life.1 -> Life.1.1
func (this *NotebookService) GetNotebookIdsAndTitles(notebookId, userId string) (cates []map[string]string) {
	i := 0
	for i < 10 { // 最多10层分类
		notebook := notebookService.GetNotebook(notebookId, userId)
		cates = append(cates, map[string]string{"Title": notebook.Title, "UrlTitle": notebook.UrlTitle, "NotebookId": notebookId})
		notebookId = notebook.ParentNotebookId.Hex()
		if len(notebookId) > 0 {
			i = i + 1
			continue
		} else {
			break
		}
	}
	for i, j := 0, len(cates)-1; i < j; i, j = i+1, j-1 { // 反转 cates
		cates[i], cates[j] = cates[j], cates[i]
	}
	return
}

// 得到某notebook
func (this *NotebookService) GetNotebook(notebookId, userId string) info.Notebook {
	notebook := info.Notebook{}
	db.GetByIdAndUserId(db.Notebooks, notebookId, userId, &notebook)
	return notebook
}
func (this *NotebookService) GetNotebookById(notebookId string) info.Notebook {
	notebook := info.Notebook{}
	db.Get(db.Notebooks, notebookId, &notebook)
	return notebook
}
func (this *NotebookService) GetNotebookByUserIdAndUrlTitle(userId, notebookIdOrUrlTitle string) info.Notebook {
	notebook := info.Notebook{}
	if IsObjectId(notebookIdOrUrlTitle) {
		db.Get(db.Notebooks, notebookIdOrUrlTitle, &notebook)
	} else {
		db.GetByQ(db.Notebooks, bson.M{"UserId": bson.ObjectIdHex(userId), "UrlTitle": encodeValue(notebookIdOrUrlTitle)}, &notebook)
	}
	return notebook
}

// 同步的方法
func (this *NotebookService) GeSyncNotebooks(userId string, afterUsn, maxEntry int) []info.Notebook {
	notebooks := []info.Notebook{}
	q := db.Notebooks.Find(bson.M{"UserId": bson.ObjectIdHex(userId), "Usn": bson.M{"$gt": afterUsn}})
	q.Sort("Usn").Limit(maxEntry).All(&notebooks)
	return notebooks
}

// 得到用户下所有的notebook 不排序 原始版本
func (this *NotebookService) GetNotebooksRaw(userId, sortField string) []info.Notebook {
	userNotebooks := []info.Notebook{}
	orQ := []bson.M{
		bson.M{"IsDeleted": false},
		bson.M{"IsDeleted": bson.M{"$exists": false}},
	}
	q := db.Notebooks.Find(bson.M{"UserId": bson.ObjectIdHex(userId), "$or": orQ})
	if sortField != "" {
		q = q.Sort(sortField)
	}
	q.All(&userNotebooks)

	if len(userNotebooks) == 0 {
		return nil
	}

	return userNotebooks
}

// 得到用户下所有的notebook
// 排序好之后返回
// [ok]
func (this *NotebookService) GetNotebooks(userId string) info.SubNotebooks {
	userNotebooks := []info.Notebook{}
	orQ := []bson.M{
		bson.M{"IsDeleted": false},
		bson.M{"IsDeleted": bson.M{"$exists": false}},
	}
	db.Notebooks.Find(bson.M{"UserId": bson.ObjectIdHex(userId), "$or": orQ}).All(&userNotebooks)

	if len(userNotebooks) == 0 {
		return nil
	}

	return ParseAndSortNotebooks(userNotebooks, true, true)
}

// share调用, 不需要删除没有父的notebook
// 不需要排序, 因为会重新排序
// 通过notebookIds得到notebooks, 并转成层次有序
func (this *NotebookService) GetNotebooksByNotebookIds(notebookIds []bson.ObjectId) info.SubNotebooks {
	userNotebooks := []info.Notebook{}
	db.Notebooks.Find(bson.M{"_id": bson.M{"$in": notebookIds}}).All(&userNotebooks)

	if len(userNotebooks) == 0 {
		return nil
	}

	return ParseAndSortNotebooks(userNotebooks, false, false)
}

// 更新父的ChildNotebookIds
func (this *NotebookService) UpdateNotebookChilds(notebookId, childNotebookId, userId, method string) (bool, info.Notebook) {
	notebook := this.GetNotebook(notebookId, userId)
	if method == "Add" {
		notebook.ChildNotebookIds = append(notebook.ChildNotebookIds, bson.ObjectIdHex(childNotebookId))
	} else if method == "Delete" {
		childs := make([]bson.ObjectId, 0, len(notebook.ChildNotebookIds))
		for _, child := range notebook.ChildNotebookIds {
			if child != bson.ObjectIdHex(childNotebookId) {
				childs = append(childs, child)
			}
		}
		notebook.ChildNotebookIds = childs
	}
	db.UpdateByIdAndUserId(db.Notebooks, notebookId, userId,
		bson.M{"$set": bson.M{"ChildNotebookIds": notebook.ChildNotebookIds}})
	return true, notebook
}

// 添加
func (this *NotebookService) AddNotebook(notebook info.Notebook) (bool, info.Notebook) {

	if notebook.NotebookId == "" {
		notebook.NotebookId = bson.NewObjectId()
	}

	userId, parentNotebookId := notebook.UserId.Hex(), notebook.ParentNotebookId.Hex()
	if len(parentNotebookId) > 0 { // 更新父的 ChildNotebookIds
		this.UpdateNotebookChilds(parentNotebookId, notebook.NotebookId.Hex(), userId, "Add")
	}

	notebook.UrlTitle = GetUrTitle(userId, notebook.Title, "notebook", notebook.NotebookId.Hex())
	notebook.Usn = userService.IncrUsn(userId)
	now := time.Now()
	notebook.CreatedTime = now
	notebook.UpdatedTime = now
	err := db.Notebooks.Insert(notebook)
	if err != nil {
		return false, notebook
	}
	return true, notebook
}

// 更新笔记, api
func (this *NotebookService) UpdateNotebookApi(userId, notebookId, title, parentNotebookId string, seq, usn int) (bool, string, info.Notebook) {
	if notebookId == "" {
		return false, "notebookIdNotExists", info.Notebook{}
	}

	// 先判断usn是否和数据库的一样, 如果不一样, 则冲突, 不保存
	notebook := this.GetNotebookById(notebookId)
	// 不存在
	if notebook.NotebookId == "" {
		return false, "notExists", notebook
	} else if notebook.Usn != usn {
		return false, "conflict", notebook
	}
	notebook.Usn = userService.IncrUsn(userId)
	notebook.Title = title

	updates := bson.M{"Title": title, "Usn": notebook.Usn, "Seq": seq, "UpdatedTime": time.Now()}
	if parentNotebookId != "" && bson.IsObjectIdHex(parentNotebookId) {
		updates["ParentNotebookId"] = bson.ObjectIdHex(parentNotebookId)
	} else {
		updates["ParentNotebookId"] = ""
	}
	ok := db.UpdateByIdAndUserIdMap(db.Notebooks, notebookId, userId, updates)
	if ok {
		return ok, "", this.GetNotebookById(notebookId)
	}
	return false, "", notebook
}

// 判断是否是blog
func (this *NotebookService) IsBlog(notebookId string) bool {
	notebook := info.Notebook{}
	db.GetByQWithFields(db.Notebooks, bson.M{"_id": bson.ObjectIdHex(notebookId)}, []string{"IsBlog"}, &notebook)
	return notebook.IsBlog
}

// 判断 Notebook 中是否含有 blog 并返回包含 blog 数量
func (this *NotebookService) HasBlog(notebookId string) int {
	return db.Count(db.Notes, bson.M{"IsBlog": true, "IsTrash": false, "IsDeleted": false,
		"Cates.NotebookId": bson.M{"$in": []string{notebookId}}})
}

// 判断是否是我的notebook
func (this *NotebookService) IsMyNotebook(notebookId, userId string) bool {
	return db.Has(db.Notebooks, bson.M{"_id": bson.ObjectIdHex(notebookId), "UserId": bson.ObjectIdHex(userId)})
}

// 更新笔记本信息
// 太广, 不用
/*
func (this *NotebookService) UpdateNotebook(notebook info.Notebook) bool {
	return db.UpdateByIdAndUserId2(db.Notebooks, notebook.NotebookId, notebook.UserId, notebook)
}
*/

// 更新笔记本标题
// [ok]
func (this *NotebookService) UpdateNotebookTitle(notebookId, userId, title string) bool {
	usn := userService.IncrUsn(userId)
	db.UpdateByQField(db.Notes, bson.M{"IsDeleted": false, "IsTrash": false, "Cates.NotebookId": notebookId}, "Cates.$.Title", title) // 修改Notes.$.Cates.$.Title
	return db.UpdateByIdAndUserIdMap(db.Notebooks, notebookId, userId, bson.M{"Title": title, "Usn": usn})
}

// 更新notebook
func (this *NotebookService) UpdateNotebook(userId, notebookId string, needUpdate bson.M) bool {
	needUpdate["UpdatedTime"] = time.Now()
	needUpdate["Usn"] = userService.IncrUsn(userId)
	return db.UpdateByIdAndUserIdMap(db.Notebooks, notebookId, userId, needUpdate)
}

// ToBlog or Not
func (this *NotebookService) ToBlog(userId, notebookId string, isBlog bool) bool {
	updates := bson.M{"IsBlog": isBlog, "Usn": userService.IncrUsn(userId)}
	// 笔记本
	db.UpdateByIdAndUserIdMap(db.Notebooks, notebookId, userId, updates)

	// 更新笔记
	q := bson.M{"UserId": bson.ObjectIdHex(userId),
		"NotebookId": bson.ObjectIdHex(notebookId)}
	data := bson.M{"IsBlog": isBlog}
	if isBlog {
		data["PublicTime"] = time.Now()
	} else {
		data["HasSelfDefined"] = false
	}
	// usn
	data["Usn"] = userService.IncrUsn(userId)
	db.UpdateByQMap(db.Notes, q, data)

	// noteContents也更新, 这个就麻烦了, noteContents表没有NotebookId
	// 先查该notebook下所有notes, 得到id
	notes := []info.Note{}
	db.ListByQWithFields(db.Notes, q, []string{"_id"}, &notes)
	if len(notes) > 0 {
		noteIds := make([]bson.ObjectId, len(notes))
		for i, each := range notes {
			noteIds[i] = each.NoteId
		}
		db.UpdateByQMap(db.NoteContents, bson.M{"_id": bson.M{"$in": noteIds}}, bson.M{"IsBlog": isBlog})
	}

	// 重新计算tags
	go (func() {
		blogService.ReCountBlogTags(userId)
	})()

	return true
}

// 查看是否有子notebook
// 先查看该notebookId下是否有notes, 没有则删除
func (this *NotebookService) DeleteNotebook(userId, notebookId string) (bool, string) {
	notebookid := bson.ObjectIdHex(notebookId)
	if db.Count(db.Notebooks, bson.M{
		"ParentNotebookId": notebookid,
		"UserId":           bson.ObjectIdHex(userId),
		"IsDeleted":        false,
	}) == 0 { // 无
		if db.Count(db.Notes, bson.M{"NotebookId": notebookid,
			"UserId":    bson.ObjectIdHex(userId),
			"IsTrash":   false,
			"IsDeleted": false}) == 0 { // 不包含trash
			notebook := this.GetNotebook(notebookId, userId)
			if len(notebook.ParentNotebookId.Hex()) > 0 { // 更新父的 ChildNotebookIds
				this.UpdateNotebookChilds(notebook.ParentNotebookId.Hex(), notebookId, userId, "Delete")
			}
			// 删除空笔记本
			return db.DeleteByIdAndUserId(db.Notebooks, notebookId, userId), ""
		}
		return false, "笔记本下有笔记"
	} else {
		return false, "笔记本下有子笔记本"
	}
}

// API调用, 删除笔记本, 不作笔记控制
func (this *NotebookService) DeleteNotebookForce(userId, notebookId string, usn int) (bool, string) {
	notebook := this.GetNotebook(notebookId, userId)
	// 不存在
	if notebook.NotebookId == "" {
		return false, "notExists"
	} else if notebook.Usn != usn {
		return false, "conflict"
	}
	if len(notebook.ParentNotebookId.Hex()) > 0 { // 更新父的 ChildNotebookIds
		this.UpdateNotebookChilds(notebook.ParentNotebookId.Hex(), notebookId, userId, "Delete")
	}
	return db.DeleteByIdAndUserId(db.Notebooks, notebookId, userId), ""
}

// 排序
// 传入 notebookId => Seq
// 为什么要传入userId, 防止修改其它用户的信息 (恶意)
// [ok]
func (this *NotebookService) SortNotebooks(userId string, notebookId2Seqs map[string]int) bool {
	if len(notebookId2Seqs) == 0 {
		return false
	}

	for notebookId, seq := range notebookId2Seqs {
		if !db.UpdateByIdAndUserIdMap(db.Notebooks, notebookId, userId, bson.M{"Seq": seq, "Usn": userService.IncrUsn(userId)}) {
			return false
		}
	}

	return true
}

// 排序和设置父
func (this *NotebookService) DragNotebooks(userId string, curNotebookId string, parentNotebookId string, siblings []string) bool {
	ok := false

	curNotebook := this.GetNotebook(curNotebookId, userId)
	if len(curNotebook.ParentNotebookId.Hex()) > 0 { // 旧的父 减
		this.UpdateNotebookChilds(curNotebook.ParentNotebookId.Hex(), curNotebookId, userId, "Delete")
	}

	// 如果没parentNotebookId, 则parentNotebookId设空
	if parentNotebookId == "" {
		ok = db.UpdateByIdAndUserIdMap(db.Notebooks, curNotebookId, userId, bson.M{"ParentNotebookId": "", "Usn": userService.IncrUsn(userId)})
	} else {
		this.UpdateNotebookChilds(parentNotebookId, curNotebookId, userId, "Add") // 新的父 加
		ok = db.UpdateByIdAndUserIdMap(db.Notebooks, curNotebookId, userId, bson.M{"ParentNotebookId": bson.ObjectIdHex(parentNotebookId), "Usn": userService.IncrUsn(userId)})
	}

	if !ok {
		return false
	}

	// 排序
	for seq, notebookId := range siblings {
		if !db.UpdateByIdAndUserIdMap(db.Notebooks, notebookId, userId, bson.M{"Seq": seq, "Usn": userService.IncrUsn(userId)}) {
			return false
		}
	}

	return true
}

// 重新统计笔记本下的笔记数目
// noteSevice: AddNote, CopyNote, CopySharedNote, MoveNote
// trashService: DeleteNote (recove不用, 都统一在MoveNote里了)
func (this *NotebookService) ReCountNotebookNumberNotes(notebookId string) bool {
	notebookIdO := bson.ObjectIdHex(notebookId)
	count := db.Count(db.Notes, bson.M{"NotebookId": notebookIdO, "IsTrash": false, "IsDeleted": false})
	// Log(count)
	// Log(notebookId)
	return db.UpdateByQField(db.Notebooks, bson.M{"_id": notebookIdO}, "NumberNotes", count)
}

func (this *NotebookService) ReCountAll() {
	/*
		// 得到所有笔记本
		notebooks := []info.Notebook{}
		db.ListByQWithFields(db.Notebooks, bson.M{}, []string{"NotebookId"}, &notebooks)

		for _, each := range notebooks {
			this.ReCountNotebookNumberNotes(each.NotebookId.Hex())
		}
	*/
}
