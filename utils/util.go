package utils

import "meetingBooking/repository/db/dao"

func BuildCategoryTree(categoryList []*dao.CategoryList) []*dao.CategoryList {
	categoryMap := make(map[uint]*dao.CategoryList)
	var roots []*dao.CategoryList

	for _, category := range categoryList {
		categoryMap[category.CategoryId] = category
		if category.ParentCategoryId == 0 {
			roots = append(roots, category)
		} else {
			parent, ok := categoryMap[category.ParentCategoryId]
			if !ok {
				parent = &dao.CategoryList{Children: []*dao.CategoryList{}}
				categoryMap[category.ParentCategoryId] = parent
			}
			parent.Children = append(parent.Children, category)
		}
	}
	return roots
}
