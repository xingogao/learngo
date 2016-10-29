package main

import (
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"os"
	"log"
	"github.com/axgle/mahonia"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"flag"
	"fmt"
)

var login *string = flag.String("login", "zemr/zemr@temr", "Use -login <login_string>")

var WhiteSpace string = string([]rune{'\t', '\n', '\v', '\f', '\r', 0x0, 0x20, 0x85, 0xA0, 0x3000})

func main() {
	flag.Parse()

	// 为log添加短文件名,方便查看行数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//log.Println("Oracle Data Mining Example")

	os.Setenv("NLS_LANG", "SIMPLIFIED CHINESE_CHINA.ZHS16GBK")

	// 用户名/密码@实例名  跟sqlplus的conn命令类似
	db, err := sql.Open("oci8", *login)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	enc := mahonia.NewEncoder("GBK")
	src := "SELECT c.id, c.content FROM zemr_emr_content c WHERE rownum < 2 AND c.title = '住院患者入院护理评估记录单'"
	strSql := enc.ConvertString(src)
	rows, err := db.Query(strSql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	dec := mahonia.NewDecoder("UTF-16")
	stmtAdd, err := db.Prepare("insert into ZEMR.NEO_NURSING_ASSESSMENT (id, emr_id, " +
		"name, sex, age, dept, bed, event, " +
		"ryzd, xyzd, lxr, " +
		"ryfs, yw, sw, qt, jws, " +
		"tw, mb, hx, xy, tz, " +
		"sz, qxzt, ms, mx, " +
		"sl, tl, kqnm, yc, " +
		"shz, st, gtfs, ljnl, " +
		"sy, syu, ysxg, sh, pf, " +
		"xb, db, lzdg, dgsj, " +
		"sm, ywfz, " +
		"ttpf, zlnl, ycpf, ddpf, gdht, " +
		"zs, " +
		"zkhl, " +
		"rq, hsqm, zkqm, " +
		"gms, " +
		"ryxj, jkzd, ph, " +
		"cl, ysj, fhfd, " +
		"yx, tzbh, " +
		"qghl, dghl, dbtc) " +
		"values " +
		"(SEQ_N_N_A_ID.NEXTVAL, ?, " +
		"?, ?, ?, ?, ?, ?, " +
		"?, ?, ?, " +
		"?, ?, ?, ?, ?, " +
		"?, ?, ?, ?, ?, " +
		"?, ?, ?, ?, " +
		"?, ?, ?, ?, " +
		"?, ?, ?, ?, " +
		"?, ?, ?, ?, ?, " +
		"?, ?, ?, ?, " +
		"?, ?, " +
		"?, ?, ?, ?, ?, " +
		"?, " +
		"?, " +
		"to_date(?, 'yyyy-MM-dd'), ?, ?, " +
		"?, " +
		"?, ?, ?, ?," +
		" ?, ?, ?, " +
		"?, ?, " +
		"?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmtAdd.Close()

	for rows.Next() {
		var emrID string
		var content []byte
		if err := rows.Scan(&emrID, &content); err != nil {
			continue
		}
		content = bytes.Trim(content, WhiteSpace)
		html := dec.ConvertString(string(content))
		row := map[string]string{"id":emrID, "content":html}
		if data, err := resolveContent(row); err != nil {
			log.Fatal(err)
		} else {
			fmt.Print(data)
			/*
			_, err := stmtAdd.Exec(emrID,
				data["姓名"], data["性别"], data["年龄"], data["科室"], data["床号"], data["住院号"],
				data["入院诊断:中医/证型"], data["西医"], data["联系人及号码"],
				data["入院方式:"], data["药物"], data["食物"], data["其他"], data["既往史:"],
				data["T"], data["P"], data["R"], data["Bp"], data["体重"],
				data["神志:"], data["情绪状态:"], data["面色:"], data["脉象:"],
				data["视力:"], data["听力:"], data["口腔黏膜:"], data["义齿:"],
				data["舌质:"], data["舌苔:"], data["沟通方式:"], data["理解能力:"],
				data["声音:"], data["食欲:"], data["饮食习惯:"], data["皮肤:"],
				data["排泄:小便:"], data["大便:"], data["留置导尿:"], data["带管时间："],
				data["睡眠:"], data["药物辅助:"],
				data["疼痛评分"], data["自理能力"], data["压疮评分"], data["跌倒评分"], data["管道滑脱评分"],
				data["二、主诉及主要阳性体征"],
				data["四、专科护理及辩证施护"],
				data["日期"], data["护士签名"], data["质控签名"],
				data["过敏史："],
				data["入院宣教"], data["健康指导"], data["陪护"],
				data["床栏"], data["约束具"], data["防滑防跌"],
				data["观察药物疗效及不良反应"], data["注意神志、瞳孔、生命体征变化"],
				data["气管切开护理"], data["导管护理"], data["保持大便通畅"])
			if err != nil {
				log.Fatal(err)
			}
			*/
		}
	}
}

func resolveContent(row map[string]string) (data map[string]string, err error) {
	data = make(map[string]string)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(row["content"]))
	if err != nil {
		return data, err
	}
	fmt.Printf("=====%v=====\r\n", row["id"])
	fmt.Print("----- td.UnderLine -----\r\n")
	doc.Find("td.UnderLine").Each(func(i int, s *goquery.Selection) {
		label := strings.Trim(s.Prev().Text(), WhiteSpace)
		if label == "" {
			label = strings.Trim(s.ParentsFiltered("p").Prev().Find("strong").Text(), WhiteSpace)
		}
		text := strings.Trim(s.Text(), WhiteSpace)
		fmt.Printf("%v\t%v\t%v\r\n", i, label, text)
		data[label] = text
	})
	value := map[string]string{"√":"1", "□":"0"}
	fmt.Print("----- td[select][class!='UnderLine'] -----\r\n")
	doc.Find("td[select][class!='UnderLine']").Each(func(i int, s *goquery.Selection) {
		label := strings.Trim(s.Prev().Text(), WhiteSpace)
		text := s.Text()
		fmt.Printf("%v\t%v\t%v\r\n", i, label, text)
		data[label] = value[text]
	})
	fmt.Print("----- span[select] -----\r\n")
	doc.Find("span[select]").Each(func(i int, s *goquery.Selection) {
		var label string
		switch  {
		/*
		case s.Nodes[0].PrevSibling != nil:
			label = strings.Trim(s.Nodes[0].PrevSibling.Data, WhiteSpace)
		*/
		case s.Nodes[0].NextSibling != nil:
			label = strings.Trim(s.Nodes[0].NextSibling.Data, WhiteSpace)
		}
		text := s.Text()
		fmt.Printf("%v\t%v\t%v\r\n", i, label, text)
		data[label] = value[text]
	})
	return data, nil
}