package errno

import (
	"github.com/zooyer/gotpl/common/consts"
	"github.com/zooyer/miskit/errors"
)

const (
	UserInitExistError = iota + 3000
	UserInitCreateError

	LicenseInvalidLicense
	LicenseWriteError
	LicenseUnAuthorize
	LicenseParseLicenseError

	SessionLoginPasswordError
	SessionLoginSaveError
	SessionLogoutSaveError

	FileGetError
	FileCountError
	FileCreateError

	ProductGetError
	ProductShowError
	ProductListError
	ProductCreateError
	ProductUpdateError
	ProductDeleteError
	ProductUpError
	ProductDownError

	ConsumerReadImportFileError
	ConsumerParseImportFileError
	ConsumerParseImportFormError
	ConsumerParseImportTimeError
	ConsumerExportWriteHeaderError
	ConsumerExportWriteDataError
	ConsumerListError
	ConsumerCreateError
	ConsumerBatchCreateError
	ConsumerBatchCreatePhoneError
	ConsumerExportError
	ConsumerExportsError
	ConsumerDeleteError
	ConsumerDeletesError

	UserHasError
	UserGetError
	UserListError
	UserCreateError
	UserUpdateError
	UserDeleteError
	UserDeleteCountError
	UserDeleteOnlyOne
	UserEnableError
	UserDisableCountError
	UserDisableOnlyOne
	UserDisableError

	SettingGetError
	SettingMGetError
	SettingSetError
	SettingMSetError
	SettingDelError
	SettingMDelError
	SettingListError
	SettingClearError

	TalentGetError
	TalentListError
	TalentSimilarError
	TalentCreateError
	TalentUpdateError
	TalentDeleteError
	TalentPublishError
	TalentFinishError

	TestGetError
	TestListError
	TestCreateError
	TestUpdateError
	TestDeleteError
)

var msg = map[int]string{
	UserInitExistError:  "用户已存在，初始化错误！！",
	UserInitCreateError: "用户初始化错误！",

	LicenseInvalidLicense:    "无效的授权码！",
	LicenseWriteError:        "授权保存错误！",
	LicenseUnAuthorize:       "未授权！",
	LicenseParseLicenseError: "授权码解析错误！",

	SessionLoginPasswordError: "登录密码错误！",
	SessionLoginSaveError:     "登录会话错误！",
	SessionLogoutSaveError:    "登出会话错误！",

	FileGetError:    "文件获取错误！",
	FileCountError:  "文件数量错误！",
	FileCreateError: "文件创建错误！",

	ProductGetError:    "产品获取错误！",
	ProductShowError:   "产品展示错误！",
	ProductListError:   "产品列表错误！",
	ProductCreateError: "产品创建错误！",
	ProductUpdateError: "产品更新错误！",
	ProductDeleteError: "产品删除错误！",
	ProductUpError:     "产品上架错误！",
	ProductDownError:   "产品下架错误！",

	ConsumerReadImportFileError:    "客户导入文件读取错误！",
	ConsumerParseImportFileError:   "客户导入文件解析错误！",
	ConsumerParseImportFormError:   "客户导入参数解析错误！",
	ConsumerParseImportTimeError:   "客户导入时间解析错误！",
	ConsumerExportWriteHeaderError: "客户导出表头错误！",
	ConsumerExportWriteDataError:   "客户导出数据错误！",
	ConsumerListError:              "客户列表错误！",
	ConsumerCreateError:            "客户创建错误！",
	ConsumerBatchCreateError:       "客户批量创建错误！",
	ConsumerBatchCreatePhoneError:  "客户批量创建手机号错误！",
	ConsumerExportError:            "客户导出错误！",
	ConsumerExportsError:           "客户批量导出错误！",
	ConsumerDeleteError:            "客户删除错误！",
	ConsumerDeletesError:           "客户批量删除错误！",

	UserHasError:          "用户已存在！",
	UserGetError:          "用户获取错误！",
	UserListError:         "用户列表错误！",
	UserCreateError:       "用户创建错误！",
	UserUpdateError:       "用户更新错误！",
	UserDeleteError:       "用户删除错误！",
	UserDeleteCountError:  "用户删除错误，获取用户数量错误！",
	UserDeleteOnlyOne:     "用户删除错误，至少保留一个用户！",
	UserEnableError:       "用户启用错误！",
	UserDisableCountError: "用户禁用错误，获取用户数量错误！",
	UserDisableOnlyOne:    "用户禁用错误，至少保留一个用户！",
	UserDisableError:      "用户禁用错误！",

	SettingGetError:   "设置获取错误！",
	SettingMGetError:  "设置批量获取错误！",
	SettingSetError:   "设置保存错误！",
	SettingMSetError:  "设置批量保存错误！",
	SettingDelError:   "设置删除错误！",
	SettingMDelError:  "设置批量删除错误！",
	SettingListError:  "设置列表错误！",
	SettingClearError: "设置清空错误！",

	TalentGetError:     "招聘获取错误！",
	TalentListError:    "招聘列表错误！",
	TalentSimilarError: "招聘相似职位错误！",
	TalentCreateError:  "招聘创建错误！",
	TalentUpdateError:  "招聘更新错误！",
	TalentDeleteError:  "招聘删除错误！",
	TalentPublishError: "招聘发布错误！",
	TalentFinishError:  "招聘终止错误！",

	TestGetError:    "测试获取错误！",
	TestListError:   "测试列表错误！",
	TestCreateError: "测试创建错误！",
	TestUpdateError: "测试更新错误！",
	TestDeleteError: "测试删除错误！",
}

func init() {
	errors.Register(consts.Name, msg)
}
