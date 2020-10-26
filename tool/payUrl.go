package tool

import (
	"bufio"
	"fmt"
	"log"
	"main/context"
	"os"
	"strings"
)

var (
	name   string
	code   string
	key    string
	secret string //商户号
	types  string //支付类型

	sql string
	url string
)

func NewUrl(app *context.Context) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("-------------------开始添加三方支付账户信息-------------------")


	fmt.Print("请输入三方支付名称(XX支付)：")
	input.Scan()
	name = input.Text()

	fmt.Print("请输入三方支付code(XXZFPay)：")
	input.Scan()
	code = input.Text()

	fmt.Print("请输入三方支付商户号：")
	input.Scan()
	key = input.Text()

	fmt.Print("请输入三方支付支付秘钥：")
	input.Scan()
	secret = input.Text()

	fmt.Print("请输入接入支付类型(1,2,3,4,5,6,7,8,9,10,11)：")
	input.Scan()
	types = input.Text()
	typesJson := "[" + types + "]"

	//第三方支付通道
	sql = fmt.Sprintf("INSERT INTO %s (name,types,code) VALUE ('%s','%s','%s')", "pay_third_party_channel", name, typesJson, code)
	row, sqlErr := app.DB2Write.Exec(sql)
	if sqlErr != nil {
		log.Fatalln(sqlErr)
	}
	row.RowsAffected()

	////第三方支付通道账号
	sql = fmt.Sprintf("INSERT INTO %s (group_id,trilateral_ident,name,merchant_key,merchant_secret_key) VALUE (%d,'%s','%s','%s','%s')", "pay_third_party_accounts", 2, code, name, key, secret)
	row2, sqlErr := app.DB2Write.Exec(sql)
	if sqlErr != nil{
		log.Fatalln(sqlErr)
	}
	accountId,_ := row2.LastInsertId()

	//匿名函数
	addPay := func(payType int, nameRight string, currentType int) {
		sql = fmt.Sprintf("INSERT INTO %s (type,`name`,grade,vip,`status`,`describe`,current_rotation_account,current_rotation_type) VALUES(%d,'%s','%s','%s',%d,'%s',%d,%d)",
			"pay_third_party_list", payType, name, "[1,2,3,5,6,7]", "[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,39]", 1, name, accountId, currentType)
		row, sqlErr := app.DB2Write.Exec(sql)
		if sqlErr != nil{
			fmt.Println(sql)
			log.Fatalln(sqlErr)
		}
		listId, _ := row.LastInsertId()

		sql = fmt.Sprintf("INSERT INTO %s (account_id,`name`,type_id,recharge_min,recharge_max,is_free,recharge_limit,quick_mode) VALUE (%d,'%s',%d,%d,%d,%d,%d,'%s')",
			"pay_third_party_account_config",accountId, name+nameRight, currentType, 0, -1, 1, 9999999999, "[]")
		row2, sqlErr := app.DB2Write.Exec(sql)
		if sqlErr != nil{
			log.Fatalln(sqlErr)
		}
		row2.RowsAffected()

		fmt.Println( name, nameRight, ":",listId)
		url = fmt.Sprintf("http://pay.com/index.php/index/pay/payList?userId=974833&money=30000&trilateralIdent=%s&id=%d&currentType=%d&accountId=%d&typeId=%d&pageId=%d&time=1577762062000&sign=c01865b88fa54ce8d27b62b1f3b96abe",
			code, listId, currentType, accountId, payType, listId)
		fmt.Println("--请求连接：", url)
	}

	typeMap := strings.Split(types, ",")
	for _, typeV := range typeMap {
		switch typeV {
		case "1":
			addPay(6, "-银行卡快捷支付", 1)
		case "2":
			addPay(7, "-云闪付扫码", 2)
		case "3":
			addPay(6, "-银行卡网银支付", 3)
		case "4":
			addPay(5, "-微信直连", 4)
		case "5":
			addPay(5, "-微信扫码", 5)
		case "6":
			addPay(4, "-支付宝直连", 6)
		case "7":
			addPay(4, "-支付宝扫码", 7)
		case "8":
			addPay(8, "-京东直连", 8)
		case "9":
			addPay(8, "-京东扫码", 9)
		case "10":
			addPay(9, "-QQ直连", 10)
		case "11":
			addPay(9, "-QQ扫码", 11)
		}
	}

}
