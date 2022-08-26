package helpers

import "client/models"

func FindMaxId(uArr []models.User) (id int) {
	if len(uArr) == 0 {
		return
	}
	if len(uArr) == 1 {
		id = uArr[0].Id
		return
	}
	for i := 0; i < len(uArr)-1; i++ {
		if uArr[i].Id > uArr[i+1.].Id {
			uArr[i], uArr[i+1] = uArr[i+1], uArr[i]
		}
	}
	id = uArr[len(uArr)-1].Id
	return
}

func FindMinId(uArr []models.User) (id int) {
	if len(uArr) == 0 {
		return
	}
	if len(uArr) == 1 {
		id = uArr[0].Id
		return
	}
	for i := 0; i < len(uArr)-1; i++ {
		if uArr[i].Id < uArr[i+1.].Id {
			uArr[i], uArr[i+1] = uArr[i+1], uArr[i]
		}
	}
	id = uArr[len(uArr)-1].Id
	return
}
