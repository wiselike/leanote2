package admin

import (
	"fmt"
	"strings"

	"github.com/wiselike/revel"

	"github.com/wiselike/leanote-of-unofficial/app/info"
	. "github.com/wiselike/leanote-of-unofficial/app/lea"
)

// admin 首页

type AdminSetting struct {
	AdminBaseController
}

// email配置
func (c AdminSetting) Email() revel.Result {
	return nil
}

// blog标签设置
func (c AdminSetting) Blog() revel.Result {
	recommendTags := configService.GetGlobalArrayConfig("recommendTags")
	newTags := configService.GetGlobalArrayConfig("newTags")
	c.ViewArgs["recommendTags"] = strings.Join(recommendTags, ",")
	c.ViewArgs["newTags"] = strings.Join(newTags, ",")
	return c.RenderTemplate("admin/setting/blog.html")
}
func (c AdminSetting) DoBlogTag(recommendTags, newTags string) revel.Result {
	re := info.NewRe()

	re.Ok = configService.UpdateGlobalArrayConfig(c.GetUserId(), "recommendTags", strings.Split(recommendTags, ","))
	re.Ok = configService.UpdateGlobalArrayConfig(c.GetUserId(), "newTags", strings.Split(newTags, ","))

	return c.RenderJSON(re)
}

// 共享设置
func (c AdminSetting) ShareNote(registerSharedUserId string,
	registerSharedNotebookPerms, registerSharedNotePerms []int,
	registerSharedNotebookIds, registerSharedNoteIds, registerCopyNoteIds []string) revel.Result {

	re := info.NewRe()
	re.Ok, re.Msg = configService.UpdateShareNoteConfig(registerSharedUserId, registerSharedNotebookPerms, registerSharedNotePerms, registerSharedNotebookIds, registerSharedNoteIds, registerCopyNoteIds)
	return c.RenderJSON(re)
}

// demo
// blog标签设置
func (c AdminSetting) Demo() revel.Result {
	c.ViewArgs["demoUsername"] = configService.GetGlobalStringConfig("demoUsername")
	c.ViewArgs["demoPassword"] = configService.GetGlobalStringConfig("demoPassword")
	return c.RenderTemplate("admin/setting/demo.html")
}
func (c AdminSetting) DoDemo(demoUsername, demoPassword string) revel.Result {
	re := info.NewRe()

	userInfo, err := authService.Login(demoUsername, demoPassword)
	if err != nil {
		re.Msg = err.Error()
		re.Ok = false
		return c.RenderJSON(re)
	}
	if userInfo.UserId == "" {
		re.Msg = "The User is Not Exists"
		return c.RenderJSON(re)
	}

	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "demoUserId", userInfo.UserId.Hex())
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "demoUsername", demoUsername)
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "demoPassword", demoPassword)

	return c.RenderJSON(re)
}

func (c AdminSetting) ExportPdf(path string) revel.Result {
	re := info.NewRe()
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "exportPdfBinPath", path)
	return c.RenderJSON(re)
}

func (c AdminSetting) DoSiteUrl(siteUrl string) revel.Result {
	re := info.NewRe()
	ok1 := configService.UpdateGlobalStringConfig(c.GetUserId(), "siteUrl", siteUrl)
	ok2 := UpdateConfig("site.url", siteUrl) == nil

	switch {
	case ok1 && ok2:
		re.Msg = "设置已生效，并且重启后也生效"
	case !ok1 && ok2:
		re.Msg = "设置未生效，需重启后才生效"
	case ok1 && !ok2:
		re.Msg = "设置已生效，重启后失效"
	case !ok1 && !ok2:
		re.Msg = "系统异常，无法修改Site's URL"
	}

	re.Ok = false // 强制弹出“提示”
	return c.RenderJSON(re)
}

// SubDomain
func (c AdminSetting) SubDomain() revel.Result {
	c.ViewArgs["str"] = configService.GlobalStringConfigs
	c.ViewArgs["arr"] = configService.GlobalArrayConfigs

	c.ViewArgs["noteSubDomain"] = configService.GetGlobalStringConfig("noteSubDomain")
	c.ViewArgs["blogSubDomain"] = configService.GetGlobalStringConfig("blogSubDomain")
	c.ViewArgs["leaSubDomain"] = configService.GetGlobalStringConfig("leaSubDomain")

	return c.RenderTemplate("admin/setting/subDomain.html")
}
func (c AdminSetting) DoSubDomain(noteSubDomain, blogSubDomain, leaSubDomain, blackSubDomains, allowCustomDomain, blackCustomDomains string) revel.Result {
	re := info.NewRe()
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "noteSubDomain", noteSubDomain)
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "blogSubDomain", blogSubDomain)
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "leaSubDomain", leaSubDomain)

	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "allowCustomDomain", allowCustomDomain)
	re.Ok = configService.UpdateGlobalArrayConfig(c.GetUserId(), "blackSubDomains", strings.Split(blackSubDomains, ","))
	re.Ok = configService.UpdateGlobalArrayConfig(c.GetUserId(), "blackCustomDomains", strings.Split(blackCustomDomains, ","))

	return c.RenderJSON(re)
}

func (c AdminSetting) OpenRegister(openRegister string) revel.Result {
	re := info.NewRe()
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "openRegister", openRegister)
	return c.RenderJSON(re)
}

func (c AdminSetting) HomePage(homePage string) revel.Result {
	re := info.NewRe()
	if homePage == "0" {
		homePage = ""
	}
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "homePage", homePage)
	return c.RenderJSON(re)
}

func (c AdminSetting) Mongodb(mongodumpPath, mongorestorePath string) revel.Result {
	re := info.NewRe()
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "mongodumpPath", mongodumpPath)
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "mongorestorePath", mongorestorePath)

	return c.RenderJSON(re)
}

func (c AdminSetting) UploadSize(uploadImageSize, uploadAvatarSize, uploadBlogLogoSize, uploadAttachSize float64) revel.Result {
	re := info.NewRe()
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "uploadImageSize", fmt.Sprintf("%v", uploadImageSize))
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "uploadAvatarSize", fmt.Sprintf("%v", uploadAvatarSize))
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "uploadBlogLogoSize", fmt.Sprintf("%v", uploadBlogLogoSize))
	re.Ok = configService.UpdateGlobalStringConfig(c.GetUserId(), "uploadAttachSize", fmt.Sprintf("%v", uploadAttachSize))
	return c.RenderJSON(re)
}
