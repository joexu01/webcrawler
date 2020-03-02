package model

import "encoding/json"

type Profile struct {
	Name     string
	Age      int
	Height   int
	Location string
	Income   string
}

func FromJsonObjToProfile(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}

type TeacherProfile struct {
	Name          string
	Gender        string
	GraduatedFrom string
	Position      string
	Education     string // 学历
	Degree        string // 学位
	IsIncumbent   string // 在职信息
	School        string // 所在单位
	Discipline    string // 学科
	Location      string // 办公地点
	Email         string
	PersonalUrl   string
}

func FromJsonObjToTeacherProfile(o interface{}) (TeacherProfile, error) {
	var profile TeacherProfile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
