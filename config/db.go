package config

import (
	"github.com/whatsauth/whatsauth"
	"os"

	"github.com/aiteung/atdb"
)

var IteungIPAddress string = os.Getenv("ITEUNGBEV1")

var MongoString string = os.Getenv("MONGOSTRRTM")

var MariaStringAkademik string = os.Getenv("MARIASTRINGAKADEMIK")

var DBUlbimariainfo = atdb.DBInfo{
	DBString: MariaStringAkademik,
	DBName:   "v6flm8jlih0151mt",
}

var DBMongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "rtmdata",
}

var MongoConn = atdb.MongoConnect(DBMongoInfo)

var Ulbimariaconn = atdb.MariaConnect(DBUlbimariainfo)

var Usertables = [4]whatsauth.LoginInfo{mhs, dosen, user, user1}

var mhs = whatsauth.LoginInfo{
	Userid:   "MhswID",
	Password: "Password",
	Phone:    "Telepon",
	Username: "Login",
	Uuid:     "simak_mst_mahasiswa",
	Login:    "2md5",
}

var dosen = whatsauth.LoginInfo{
	Userid:   "NIDN",
	Password: "Password",
	Phone:    "Handphone",
	Username: "Login",
	Uuid:     "simak_mst_dosen",
	Login:    "2md5",
}

var user = whatsauth.LoginInfo{
	Userid:   "user_id",
	Password: "user_password",
	Phone:    "phone",
	Username: "user_name",
	Uuid:     "simak_besan_users",
	Login:    "2md5",
}

var user1 = whatsauth.LoginInfo{
	Userid:   "user_id",
	Password: "user_password",
	Phone:    "user_phone",
	Username: "user_name",
	Uuid:     "besan_users",
	Login:    "2md5",
}
