package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

var WhiteSpace string = string([]rune{'\t', '\n', '\v', '\f', '\r', 0x0, 0x20, 0x85, 0xA0, 0x3000})

func main() {
	content := `
<HTML XMLNS:CUSTOMTAG><HEAD><TITLE>无标题文档</TITLE>
<STYLE>
BODY {
	FONT-SIZE: 14px; FONT-FAMILY: 宋体
}
TBODY {
	FONT-SIZE: 14px; FONT-FAMILY: 宋体
}
P {
	MARGIN-TOP: 0px; MARGIN-BOTTOM: 0px
}
.border_td {
	BORDER-RIGHT: #000000 1px solid; BORDER-TOP: #000000 1px solid; BORDER-LEFT: #000000 1px solid; BORDER-BOTTOM: #000000 1px solid; BACKGROUND-COLOR: #ffffff
}
.no_border_td {
	BACKGROUND-COLOR: #ffffff
}
.border_table {
	BORDER-RIGHT: medium none; BORDER-TOP: medium none; BORDER-LEFT: medium none; BORDER-BOTTOM: medium none; BORDER-COLLAPSE: collapse; mso-border-alt: solid windowtext .5pt; mso-padding-alt: 0cm 5.4pt 0cm 5.4pt
}
.no_border_table {
	BACKGROUND-COLOR: darkgray
}
</STYLE>

<STYLE>
BODY {
	FONT-SIZE: 14px; FONT-FAMILY: 宋体
}
TBODY {
	FONT-SIZE: 14px; FONT-FAMILY: 宋体
}
P {
	MARGIN-TOP: 0px; MARGIN-BOTTOM: 0px
}
.border_td {
	BORDER-RIGHT: #000000 1px solid; BORDER-TOP: #000000 1px solid; BORDER-LEFT: #000000 1px solid; BORDER-BOTTOM: #000000 1px solid; BACKGROUND-COLOR: #ffffff
}
.no_border_td {
	BACKGROUND-COLOR: #ffffff
}
.border_table {
	BORDER-RIGHT: medium none; BORDER-TOP: medium none; BORDER-LEFT: medium none; BORDER-BOTTOM: medium none; BORDER-COLLAPSE: collapse; mso-border-alt: solid windowtext .5pt; mso-padding-alt: 0cm 5.4pt 0cm 5.4pt
}
.no_border_table {
	BACKGROUND-COLOR: darkgray
}
</STYLE>

<STYLE>
BODY {
	FONT-SIZE: 14px; FONT-FAMILY: 宋体
}
TBODY {
	FONT-SIZE: 14px; FONT-FAMILY: 宋体
}
P {
	MARGIN-TOP: 0px; MARGIN-BOTTOM: 0px
}
.border_td {
	BORDER-RIGHT: #000000 1px solid; BORDER-TOP: #000000 1px solid; BORDER-LEFT: #000000 1px solid; BORDER-BOTTOM: #000000 1px solid; BACKGROUND-COLOR: #ffffff
}
.no_border_td {
	BACKGROUND-COLOR: #ffffff
}
.border_table {
	BORDER-RIGHT: medium none; BORDER-TOP: medium none; BORDER-LEFT: medium none; BORDER-BOTTOM: medium none; BORDER-COLLAPSE: collapse; mso-border-alt: solid windowtext .5pt; mso-padding-alt: 0cm 5.4pt 0cm 5.4pt
}
.no_border_table {
	BACKGROUND-COLOR: darkgray
}
</STYLE>

<STYLE id=Revision>
.INSERT1 {
	COLOR: #0099ff; TEXT-DECORATION: underline
}
.DELETE1 {
	COLOR: #ff00cc; TEXT-DECORATION: line-through
}
.INSERT2 {
	COLOR: #0000ff; TEXT-DECORATION: underline
}
.DELETE2 {
	COLOR: #ff0000; TEXT-DECORATION: line-through
}
.INSERT3 {
	COLOR: #000099; TEXT-DECORATION: underline
}
.DELETE3 {
	COLOR: #aa0000; TEXT-DECORATION: line-through
}
.InsertCtrl {
	BORDER-RIGHT: #0000ff 1px solid; BORDER-TOP: #0000ff 1px solid; BORDER-LEFT: #0000ff 1px solid; BORDER-BOTTOM: #0000ff 1px solid
}
.DeleteCtrl {
	BORDER-RIGHT: #ff0000 1px solid; BORDER-TOP: #ff0000 1px solid; BORDER-LEFT: #ff0000 1px solid; BORDER-BOTTOM: #ff0000 1px solid
}
</STYLE>

<META http-equiv=Content-Type content="text/html; charset=gb2312">
<STYLE type=text/css>
<!--
TD {  font-size: 14px;FONT-FAMILY: 宋体;}
.UnderLine {
	font-size: 14px;FONT-FAMILY: 宋体;
	border-bottom:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.UnderLineX {
	font-size: 14px;FONT-FAMILY: 宋体;
	border-bottom:#CDCDCD 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.UnderLine2 {
	font-size: 14px;FONT-FAMILY: 宋体;
	border-bottom:#000000 2px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.LeftLine {
	font-size: 14px;FONT-FAMILY: 宋体;
	border-left:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.RightLine {
	font-size: 14px;FONT-FAMILY: 宋体;
	border-right:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.TopLine {
	font-size: 14px;FONT-FAMILY: 宋体;
	border-top:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.TopLine2 {
	font-size: 14px;FONT-FAMILY: 宋体;
	border-top:#000000 2px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.R_B_Line {
	font-size: 14px;FONT-FAMILY: 宋体;
	border-right:#000000 1px solid;
	border-bottom:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.R_B_Line2 {
	font-size: 14px;FONT-FAMILY: 宋体;
	border-right:#000000 1px solid;
	border-bottom:#000000 2px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
}
.L_R_B_Line {	font-size: 14px;FONT-FAMILY: 宋体;
	border-left:#000000 1px solid;
	border-right:#000000 1px solid;
	border-bottom:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.B_Line {	font-size: 14px;FONT-FAMILY: 宋体
	border-bottom:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.T_L_R_B_Line {	font-size: 14px;FONT-FAMILY: 宋体;
	border-top:#000000 1px solid;
	border-left:#000000 1px solid;
	border-right:#000000 1px solid;
	border-bottom:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
	.T_B_Line {	font-size: 14px;FONT-FAMILY: 宋体;
	border-top:#000000 1px solid;
	border-bottom:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.T_L_R_Line {	font-size: 14px;FONT-FAMILY: 宋体;
	border-top:#000000 1px solid;
	border-left:#000000 1px solid;
	border-right:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.T_R_B_Line {	font-size: 14px;FONT-FAMILY: 宋体;
	border-top:#000000 1px solid;
	border-right:#000000 1px solid;
	border-bottom:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
	.T_R_Line {	font-size: 14px;FONT-FAMILY: 宋体;
	border-top:#000000 1px solid;
	border-right:#000000 1px solid;
	borderColor=#000000
	borderColorLight=#ffffff
	borderColorDark=#000000
	}
.cBig {
	font-size: 18px; FONT-FAMILY: "宋体";
	}
.cTitle {
	font-family: "黑体";
	font:22px}
.cHspl {
	font-family: "楷体";
	font-size: 16px;
}
.style2 {font-size: 12px}
.style4 {font-size: 12px; FONT-FAMILY: 宋体; border-top: #000000 1px solid; border-right: #000000 1px solid; }
.style5 {color: #FF0000}
.border_table1 {	BORDER-RIGHT: medium none; BORDER-TOP: medium none; BORDER-LEFT: medium none; BORDER-BOTTOM: medium none; BORDER-COLLAPSE: collapse; mso-border-alt: solid windowtext .5pt; mso-padding-alt: 0cm 5.4pt 0cm 5.4pt
}
.border_table1 {	BORDER-RIGHT: medium none; BORDER-TOP: medium none; BORDER-LEFT: medium none; BORDER-BOTTOM: medium none; BORDER-COLLAPSE: collapse; mso-border-alt: solid windowtext .5pt; mso-padding-alt: 0cm 5.4pt 0cm 5.4pt
}
.border_table1 {	BORDER-RIGHT: medium none; BORDER-TOP: medium none; BORDER-LEFT: medium none; BORDER-BOTTOM: medium none; BORDER-COLLAPSE: collapse; mso-border-alt: solid windowtext .5pt; mso-padding-alt: 0cm 5.4pt 0cm 5.4pt
}
-->
</STYLE>

<META content="MSHTML 6.00.2900.6357" name=GENERATOR>
<STYLE id=system_default_css>BODY {
	TEXT-JUSTIFY: distribute; FONT-SIZE: 14px; FONT-FAMILY: 宋体; TEXT-ALIGN: justify
}
TBODY {
	TEXT-JUSTIFY: distribute; FONT-SIZE: 14px; FONT-FAMILY: 宋体; TEXT-ALIGN: justify
}
P {
	MARGIN-TOP: 0px; MARGIN-BOTTOM: 0px
}
.QUALITY1 {
	BORDER-BOTTOM: red thin dashed
}
.QUALITY2 {
	BORDER-BOTTOM: orange thin dashed
}
.QUALITY3 {
	BORDER-BOTTOM: blue thin dashed
}
</STYLE>
</HEAD>
<BODY bgColor=#ece9d8 create_dept_id="389" create_dept_name="脑病四科" reportInfection="0">
<TABLE borderColor=#000000 cellSpacing=0 cellPadding=0 width=697 align=center border=0 edit="0">
<TBODY edit="0">
<TR edit="0">
<TD style="LINE-HEIGHT: 150%" align=right width=697 edit="0">&nbsp;</TD></TR>
<TR edit="0">
<TD class=cHspl style="LINE-HEIGHT: 150%" align=middle edit="0" STYLE8>安 徽 中 医 药 大 学 第 二 附 属 医 院</TD></TR>
<TR edit="0">
<TD class=cTitle id=tt align=middle colSpan=6 height=32 edit="0" STYLE8>住 院 患 者 护 理 评 估 记 录 单</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" height=20 cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=left width=30 height=20 edit="0">姓名</TD>
<TD class=UnderLine id=STD_DE_PATIENT_NAME vAlign=center align=middle width=68 bgColor=#ffffff xid="base.name" type="0" value="patient.name" maxlength="64" prompt="服务对象标识/病人姓名" datatype="VARCHAR2">王伟</TD>
<TD vAlign=center align=left width=34 edit="0">性别</TD>
<TD class=UnderLine vAlign=center align=middle width=26 bgColor=#d6dff7 xid="base.sex" type="0" value="patient.sex" select="0/男/女">男</TD>
<TD vAlign=center align=middle width=35 edit="0">年龄</TD>
<TD class=UnderLine vAlign=center align=middle width=40 bgColor=#ffffff xid="base.age" type="0" value="patient.age" maxlength="36" datatype="VARCHAR2" nextid="patient_marriage">46岁</TD>
<TD vAlign=center align=middle width=37 edit="0">科室</TD>
<TD class=UnderLine vAlign=center align=middle width=155 bgColor=#ffffff type="1" value="residence.now.dept">脑病四科</TD>
<TD vAlign=center align=right width=38 edit="0">床号</TD>
<TD class=UnderLine vAlign=center align=middle width=70 bgColor=#ffffff>&nbsp;4床</TD>
<TD vAlign=center align=right width=53 edit="0">住院号</TD>
<TD class=UnderLine vAlign=center align=left width=114 bgColor=#ffffff type="1" value="residence.event">300000748002</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" borderColor=#424142 height=20 cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=top align=left width=133 height=20 edit="0">入院诊断:中医/证型</TD>
<TD class=UnderLine vAlign=top align=left width=198 bgColor=#ffffff>&nbsp;中风/气虚血瘀</TD>
<TD vAlign=top align=left width=40 edit="0">西医</TD>
<TD class=UnderLine vAlign=top align=left width=135 bgColor=#ffffff>&nbsp;脑梗塞</TD>
<TD vAlign=top align=left width=93 edit="0">联系人及号码</TD>
<TD class=UnderLine vAlign=top align=left width=101 bgColor=#ffffff>&nbsp;1234556</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=top align=left width=70 height=20 edit="0">入院方式:</TD>
<TD class=UnderLine vAlign=top align=middle width=62 bgColor=#d6dff7 select="0/步行/轮椅/平车/&nbsp;" code>轮椅</TD>
<TD vAlign=top align=left width=66 height=20 edit="0">&nbsp;过敏史：</TD>
<TD vAlign=top align=left width=36 bgColor=#d6dff7 select="0/无/有：/&nbsp;" code>无</TD>
<TD vAlign=top align=left width=34 edit="0">药物</TD>
<TD class=UnderLine vAlign=top align=left width=129 bgColor=#ffffff>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; /</TD>
<TD vAlign=top align=left width=36 edit="0">食物</TD>
<TD class=UnderLine vAlign=top align=left width=99 bgColor=#ffffff>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; /</TD>
<TD vAlign=top align=left width=36 edit="0">其他</TD>
<TD class=UnderLine vAlign=top align=left width=132 bgColor=#ffffff>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; /</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" height=20 cellSpacing=0 cellPadding=0 width=700 align=center bgColor=#efebde border=0 edit="0">
<TBODY edit="0">
<TR class=L9 vAlign=bottom edit="0">
<TD vAlign=center align=left width=57 height=20 edit="0">既往史:</TD>
<TD class=UnderLine vAlign=center align=left width=643 bgColor=#ffffff>&nbsp;高血压</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY edit="0">
<TR class=L9 vAlign=bottom edit="0">
<TD vAlign=center align=left width=135 height=20 edit="0"><STRONG edit="0">一、护理评估：</STRONG></TD>
<TD vAlign=center align=middle width=565 edit="0">&nbsp;</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY edit="0">
<TR class=L9 edit="0">
<TD vAlign=center borderColor=#404040 align=left width=14 height=20 edit="0">T</TD>
<TD class=UnderLine vAlign=center borderColor=#404040 align=left width=56 bgColor=#ffffff>&nbsp;36.5</TD>
<TD vAlign=center borderColor=#404040 align=left width=31 edit="0"><SPAN class=style2>℃</SPAN></TD>
<TD vAlign=center borderColor=#404040 align=left width=19 edit="0">P</TD>
<TD class=UnderLine vAlign=center borderColor=#404040 align=left width=53 bgColor=#ffffff>&nbsp;78</TD>
<TD vAlign=center borderColor=#404040 align=left width=49 edit="0">次/分</TD>
<TD vAlign=center borderColor=#404040 align=left width=21 edit="0">R</TD>
<TD class=UnderLine vAlign=center borderColor=#404040 align=left width=58 bgColor=#ffffff>&nbsp;16</TD>
<TD vAlign=center borderColor=#404040 align=left width=52 edit="0">次/分</TD>
<TD vAlign=center borderColor=#404040 align=left width=16 edit="0">Bp</TD>
<TD class=UnderLine vAlign=center borderColor=#404040 align=left width=84 bgColor=#ffffff>&nbsp;123/78</TD>
<TD vAlign=center borderColor=#404040 align=left width=72 edit="0">mmHg</TD>
<TD vAlign=center borderColor=#404040 align=left width=39 edit="0">体重</TD>
<TD class=UnderLine vAlign=center borderColor=#404040 align=left width=81 bgColor=#ffffff>&nbsp;70</TD>
<TD vAlign=center borderColor=#404040 align=left width=55 edit="0">kg</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 2px; MARGIN-BOTTOM: 0px" edit="0">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=left width=45 height=19 edit="0">神志:</TD>
<TD class=UnderLine vAlign=center align=left width=66 bgColor=#d6dff7 select="0/清楚/嗜睡/意识模糊/昏睡/浅昏迷/深昏迷/&nbsp;" code>清楚</TD>
<TD vAlign=center align=right width=71 height=19 edit="0">情绪状态:</TD>
<TD class=UnderLine vAlign=center align=middle width=88 bgColor=#d6dff7 select="0/稳定/焦虑/紧张/恐惧/&nbsp;" code>稳定</TD>
<TD vAlign=center align=right width=44 height=19 edit="0">面色:</TD>
<TD class=UnderLine vAlign=center align=left width=93 bgColor=#d6dff7 select="0/红润/少华/苍白/萎黄/晦暗/紫绀/&nbsp;" code>少华</TD>
<TD vAlign=center align=right width=47 height=19 edit="0">脉象:</TD>
<TD class=UnderLine vAlign=center align=left width=246 bgColor=#d6dff7 select="0/弦脉/滑脉/弦滑脉/细脉/沉脉/数脉/缓脉/结代脉/&nbsp;" code>弦滑脉</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 2px; MARGIN-BOTTOM: 0px" edit="0">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=left width=43 height=20 edit="0">视力:</TD>
<TD class=UnderLine vAlign=center align=left width=87 bgColor=#d6dff7 select="0/正常/异常/&nbsp;" code>正常</TD>
<TD vAlign=center align=right width=44 height=20 edit="0">听力:</TD>
<TD class=UnderLine vAlign=center align=left width=154 bgColor=#d6dff7 select="0/正常/异常/&nbsp;" code>正常</TD>
<TD vAlign=center align=right width=73 height=20 edit="0">口腔黏膜:</TD>
<TD class=UnderLine vAlign=center align=left width=156 bgColor=#d6dff7 select="0/正常/破损/&nbsp;" code>正常</TD>
<TD vAlign=center align=right width=37 height=20 edit="0">义齿:</TD>
<TD class=UnderLine vAlign=center align=left width=106 bgColor=#d6dff7 select="0/无/有/&nbsp;" code>无</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 2px; MARGIN-BOTTOM: 0px" edit="0">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=left width=43 height=20 edit="0">舌质:</TD>
<TD class=UnderLine vAlign=center align=left width=87 bgColor=#d6dff7 select="0/淡舌/淡红舌/淡白舌/红舌/绛舌/紫舌/&nbsp;" code>淡舌</TD>
<TD vAlign=center align=right width=45 height=20 edit="0">舌苔:</TD>
<TD class=UnderLine vAlign=center align=left width=73 bgColor=#d6dff7 select="0/薄白苔/黄苔/白苔/白腻苔/黄腻苔/灰苔/黑苔/少苔/nbsp;" code>白苔</TD>
<TD vAlign=center align=right width=70 height=20 edit="0">沟通方式:</TD>
<TD class=UnderLine vAlign=center align=left width=115 bgColor=#d6dff7 select="0/语言/文字/手势/&nbsp;" code>语言</TD>
<TD vAlign=center align=right width=76 height=20 edit="0">理解能力:</TD>
<TD class=UnderLine vAlign=center align=left width=191 bgColor=#d6dff7 height=20 select="0/良好/一般/差/&nbsp;" code>良好</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=left width=43 height=20 edit="0">声音:</TD>
<TD class=UnderLine vAlign=center align=left width=70 bgColor=#d6dff7 select="0/正常/嘶哑/语音低微/咳声无力或重浊/失声/&nbsp;" code>正常</TD>
<TD vAlign=center align=left width=43 height=20 edit="0">食欲:</TD>
<TD class=UnderLine vAlign=center width=105 bgColor=#d6dff7 select="0/良好/亢进/纳差/纳呆/鼻饲/禁食/&nbsp;" code>良好</TD>
<TD vAlign=center align=left width=70 height=20 edit="0">饮食习惯:</TD>
<TD class=UnderLine vAlign=center align=left width=167 bgColor=#d6dff7 select="0/咸/甜/辛辣/油腻/清淡/忌食/&nbsp;" code>清淡</TD>
<TD vAlign=center align=left width=46 height=20 edit="0">皮肤:</TD>
<TD class=UnderLine vAlign=center align=left width=156 bgColor=#d6dff7 height=20 select="0/正常/水肿/黄染/紫绀/皮疹/瘀斑/不完整/&nbsp;" code>正常</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=left width=71 height=20 edit="0">排泄:小便:</TD>
<TD class=UnderLine vAlign=center align=left width=61 bgColor=#d6dff7 select="0/正常/清长/短涩/失禁/血尿/尿潴留/保留导尿/人工瘘管/&nbsp;" code>正常</TD>
<TD vAlign=center align=left width=39 height=20 edit="0">大便:</TD>
<TD class=UnderLine vAlign=center align=left width=96 bgColor=#d6dff7 select="0/正常/便秘/便溏/失禁/血便/肠造瘘/&nbsp;" code>正常</TD>
<TD vAlign=center align=left width=67 height=20 edit="0">留置导尿:</TD>
<TD class=UnderLine vAlign=center align=left width=140 bgColor=#d6dff7 select="1/气管套管/胃管/尿管/深静脉置管/PICC置管/引流管/&nbsp;">&nbsp;/</TD>
<TD vAlign=center align=left width=73 height=20 edit="0">带管时间：</TD>
<TD class=UnderLine vAlign=center align=left width=153 bgColor=#ffccff date="yyyy-mm-dd hh:MM" dict="#date">&nbsp;</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=left width=47 height=20 edit="0">睡眠:</TD>
<TD class=UnderLine vAlign=center align=left width=65 bgColor=#d6dff7 select="0/正常/异常/&nbsp;" code>正常</TD>
<TD vAlign=center align=left width=67 height=20 edit="0">药物辅助:</TD>
<TD class=UnderLine vAlign=center align=left width=66 bgColor=#d6dff7 select="0/无/有/&nbsp;" code>无</TD>
<TD vAlign=center align=left width=455 height=20 edit="0">&nbsp;</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY edit="0">
<TR class=L9 edit="0">
<TD vAlign=center borderColor=#404040 align=left width=68 height=20 edit="0">疼痛评分</TD>
<TD class=UnderLine vAlign=center borderColor=#404040 align=left width=47 bgColor=#ffffff>&nbsp;/</TD>
<TD vAlign=center borderColor=#404040 align=left width=56 edit="0">分</TD>
<TD vAlign=center borderColor=#404040 align=left width=72 edit="0">自理能力</TD>
<TD class=UnderLine vAlign=center borderColor=#404040 align=left width=49 bgColor=#ffffff>&nbsp;</TD>
<TD vAlign=center borderColor=#404040 align=left width=33 edit="0">分</TD>
<TD vAlign=center borderColor=#404040 align=left width=58 edit="0">压疮评分</TD>
<TD class=UnderLine vAlign=center borderColor=#404040 align=left width=33 bgColor=#ffffff>&nbsp;</TD>
<TD vAlign=center borderColor=#404040 align=left width=25 edit="0">分</TD>
<TD vAlign=center borderColor=#404040 align=left width=59 edit="0">跌倒评分</TD>
<TD class=UnderLine vAlign=center borderColor=#404040 align=left width=38 bgColor=#ffffff>&nbsp;</TD>
<TD vAlign=center borderColor=#404040 align=left width=21 edit="0">分</TD>
<TD vAlign=center borderColor=#404040 align=left width=95 edit="0">管道滑脱评分</TD>
<TD class=UnderLine vAlign=center borderColor=#404040 align=left width=26 bgColor=#ffffff>&nbsp;</TD>
<TD vAlign=center borderColor=#404040 align=left width=20 edit="0">分</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=left height=20 edit="0"><STRONG edit="0">二、主诉及主要阳性体征</STRONG></TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 2px; MARGIN-BOTTOM: 0px" edit="0">
<TABLE style="WORD-BREAK: break-all; LINE-HEIGHT: 150%" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD class=UnderLine vAlign=top align=left bgColor=#ffffff height=20>&nbsp;右侧肢体活动不利半年</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 2px; MARGIN-BOTTOM: 0px" edit="0">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=left height=20 edit="0"><STRONG edit="0">三、护理常规:</STRONG></TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD style="LINE-HEIGHT: 150%" vAlign=center align=left height=20 edit="0"><SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/" code>√</SPAN>入院宣教 <SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/" code>√</SPAN>健康指导 <SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/" code>√</SPAN>陪护 <SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/" code>√</SPAN>床栏 <SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/">□</SPAN>约束具 <SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/" code>√</SPAN>防滑防跌 <SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/" code>√</SPAN>观察药物疗效及不良反应 <SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/" code>√</SPAN>注意神志、瞳孔、生命体征变化 <SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/" code>√</SPAN>气管切开护理 <SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/">□</SPAN>导管护理 <SPAN style="BACKGROUND-COLOR: #d6dff7" edit="0" select="0/√/□/" code>√</SPAN>保持大便通畅 </TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=left height=20 edit="0"><STRONG edit="0">四、专科护理及辩证施护</STRONG></TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all; LINE-HEIGHT: 150%" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD class=UnderLine vAlign=center align=left bgColor=#ffffff height=20>二级护理，低盐低脂饮食</TD></TR></TBODY></TABLE>
<P style="MARGIN-TOP: 5px; MARGIN-BOTTOM: 0px">
<TABLE style="WORD-BREAK: break-all" cellSpacing=0 cellPadding=0 width=700 align=center border=0 edit="0">
<TBODY>
<TR class=L9 vAlign=bottom>
<TD vAlign=center align=right width=159 height=20 edit="0">日期</TD>
<TD class=UnderLine vAlign=center align=left width=146 bgColor=#ffccff date="yyyy-mm-dd" dict="#date">2014-10-24</TD>
<TD vAlign=center align=right width=81 edit="0">护士签名</TD>
<TD class=UnderLine vAlign=center align=left width=95 bgColor=#f0f0f0 dict="8">&nbsp;李静</TD>
<TD vAlign=center align=right width=89 edit="0">质控签名</TD>
<TD class=UnderLine vAlign=center align=left width=130 bgColor=#f0f0f0 dict="8">&nbsp;</TD></TR></TBODY></TABLE></P></BODY></HTML>
`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return
	}
	/*
	doc.Find("[xid]").Each(func(i int, s *goquery.Selection) {
		a, _ := s.Attr("xid")
		fmt.Println(a, "\t", s.Text())
	})
	doc.Find("[value^='patient.']").Each(func(i int, s *goquery.Selection) {
		a, _ := s.Attr("value")
		fmt.Println(a, "\t", s.Text())
	})
	doc.Find("[value^='residence.']").Each(func(i int, s *goquery.Selection) {
		a, _ := s.Attr("value")
		fmt.Println(a, "\t", s.Text())
	})
	*/
	data := make(map[string]string)
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
		label := strings.Trim(s.Nodes[0].NextSibling.Data, WhiteSpace)
		text := s.Text()
		fmt.Printf("%v\t%v\t%v\r\n", i, label, text)
		data[label] = value[text]
	})
	fmt.Print(len(data), data)
}
