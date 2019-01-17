package models

import (
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
	"database/sql"
	"log"
	"time"
	"github.com/astaxie/beego/orm"
)

type MovieInfo struct {
	Id                   int64
	Movie_id             int64
	Movie_name           string
	Movie_pic            string
	Movie_director       string
	Movie_writer         string
	Movie_country        string
	Movie_language       string
	Movie_main_character string
	Movie_type           string
	Movie_on_time        string
	Movie_span           string
	Movie_grade          string
	Create_time          string
}

var db *sql.DB

func init() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	db, _ = sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/go_crawl?charset=utf8")
}

//导演名称
func GetMovieDirector(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	return string(result[0][1])

}

//电影名称
func GetMovieName(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	return string(result[0][1])
}

//主演
func GetMovieMainCharacters(movieHtml string) string {
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	mainCharacters := ""
	for _, v := range result {
		mainCharacters += v[1] + "/"
	}

	return strings.Trim(mainCharacters, "/")
}

//电影评分
func GetMovieGrade(movieHtml string) string {
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	return string(result[0][1])
}

//电影分类
func GetMovieGenre(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	movieGenre := ""
	for _, v := range result {
		movieGenre += v[1] + "/"
	}
	return strings.Trim(movieGenre, "/")
}

//上映时间
func GetMovieOnTime(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	return string(result[0][1])
}

//电影时长
func GetMovieRunningTime(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	return string(result[0][1])
}

//添加到数据库中
func AddMovie(movie_info *MovieInfo) (int64, error) {
	//id,err := db.Insert(movie_info)

	result, err := db.Exec("INSERT INTO movie_info ("+
		"id, movie_id, movie_name, movie_pic, movie_director, "+
		"movie_writer,movie_country,movie_language,movie_main_character,movie_type,"+
		"movie_on_time,movie_span,movie_grade,remark,_create_time,_modify_time,_status) "+
		"VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		movie_info.Id, movie_info.Movie_id, movie_info.Movie_name, movie_info.Movie_pic, movie_info.Movie_director,
		movie_info.Movie_writer, movie_info.Movie_country, movie_info.Movie_language, movie_info.Movie_main_character, movie_info.Movie_type,
		movie_info.Movie_on_time, movie_info.Movie_span, movie_info.Movie_grade, "", movie_info.Create_time, time.Now().Format("2006-1-2 15:04:05"), 1)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, err
}

//获取页面的URL
func GetMovieUrls(movieHtml string) []string {
	reg := regexp.MustCompile(`<a.*?href="(https://movie.douban.com/.*?)"`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	var movieSets []string
	for _, v := range result {
		movieSets = append(movieSets, v[1])
	}

	return movieSets
}
