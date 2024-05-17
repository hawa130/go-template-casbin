// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"github.com/hawa130/computility-cloud/ent/schema\",\"Package\":\"github.com/hawa130/computility-cloud/ent\",\"Schemas\":[{\"name\":\"Permission\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"roles\",\"type\":\"Role\",\"ref_name\":\"permissions\",\"inverse\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"xid.ID\",\"PkgPath\":\"github.com/rs/xid\",\"PkgName\":\"xid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":{\"Bytes\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Compare\":{\"In\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null}],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Counter\":{\"In\":[],\"Out\":[{\"Name\":\"int32\",\"Ident\":\"int32\",\"Kind\":5,\"PkgPath\":\"\",\"Methods\":null}]},\"Encode\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"IsNil\":{\"In\":[],\"Out\":[{\"Name\":\"bool\",\"Ident\":\"bool\",\"Kind\":1,\"PkgPath\":\"\",\"Methods\":null}]},\"IsZero\":{\"In\":[],\"Out\":[{\"Name\":\"bool\",\"Ident\":\"bool\",\"Kind\":1,\"PkgPath\":\"\",\"Methods\":null}]},\"Machine\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalJSON\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Pid\":{\"In\":[],\"Out\":[{\"Name\":\"uint16\",\"Ident\":\"uint16\",\"Kind\":9,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"time.Time\",\"Kind\":25,\"PkgPath\":\"time\",\"Methods\":null}]},\"UnmarshalJSON\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}]},{\"name\":\"Role\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"permissions\",\"type\":\"Permission\"},{\"name\":\"users\",\"type\":\"User\",\"ref_name\":\"roles\",\"inverse\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"xid.ID\",\"PkgPath\":\"github.com/rs/xid\",\"PkgName\":\"xid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":{\"Bytes\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Compare\":{\"In\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null}],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Counter\":{\"In\":[],\"Out\":[{\"Name\":\"int32\",\"Ident\":\"int32\",\"Kind\":5,\"PkgPath\":\"\",\"Methods\":null}]},\"Encode\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"IsNil\":{\"In\":[],\"Out\":[{\"Name\":\"bool\",\"Ident\":\"bool\",\"Kind\":1,\"PkgPath\":\"\",\"Methods\":null}]},\"IsZero\":{\"In\":[],\"Out\":[{\"Name\":\"bool\",\"Ident\":\"bool\",\"Kind\":1,\"PkgPath\":\"\",\"Methods\":null}]},\"Machine\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalJSON\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Pid\":{\"In\":[],\"Out\":[{\"Name\":\"uint16\",\"Ident\":\"uint16\",\"Kind\":9,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"time.Time\",\"Kind\":25,\"PkgPath\":\"time\",\"Methods\":null}]},\"UnmarshalJSON\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}]},{\"name\":\"User\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"roles\",\"type\":\"Role\"}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"xid.ID\",\"PkgPath\":\"github.com/rs/xid\",\"PkgName\":\"xid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":{\"Bytes\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Compare\":{\"In\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null}],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Counter\":{\"In\":[],\"Out\":[{\"Name\":\"int32\",\"Ident\":\"int32\",\"Kind\":5,\"PkgPath\":\"\",\"Methods\":null}]},\"Encode\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"IsNil\":{\"In\":[],\"Out\":[{\"Name\":\"bool\",\"Ident\":\"bool\",\"Kind\":1,\"PkgPath\":\"\",\"Methods\":null}]},\"IsZero\":{\"In\":[],\"Out\":[{\"Name\":\"bool\",\"Ident\":\"bool\",\"Kind\":1,\"PkgPath\":\"\",\"Methods\":null}]},\"Machine\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalJSON\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Pid\":{\"In\":[],\"Out\":[{\"Name\":\"uint16\",\"Ident\":\"uint16\",\"Kind\":9,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"time.Time\",\"Kind\":25,\"PkgPath\":\"time\",\"Methods\":null}]},\"UnmarshalJSON\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"annotations\":{\"EntGQL\":{\"OrderField\":\"CREATED_AT\",\"Skip\":16}}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1},\"annotations\":{\"EntGQL\":{\"OrderField\":\"UPDATED_AT\",\"Skip\":48}}},{\"name\":\"nickname\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"username\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"optional\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"email\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"phone\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"password\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"sensitive\":true}],\"indexes\":[{\"fields\":[\"created_at\"]},{\"fields\":[\"updated_at\"]},{\"unique\":true,\"fields\":[\"username\",\"email\",\"phone\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}],\"annotations\":{\"EntGQL\":{\"MutationInputs\":[{\"IsCreate\":true},{}],\"QueryField\":{},\"RelayConnection\":true}}}],\"Features\":[\"namedges\",\"privacy\",\"entql\",\"schema/snapshot\"]}"