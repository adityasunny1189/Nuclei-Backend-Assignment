package test

import (
	"nuclei-assignment-3/models"
	"nuclei-assignment-3/utils"
	"testing"
)

type tree []models.Node

var testcases = []struct {
	testCase    int
	testNode    int
	expectedVal bool
	testTree    tree
}{
	{
		testCase:    1,
		testNode:    1,
		expectedVal: true,
		testTree: tree{
			{
				NodeId:  1,
				Name:    "A",
				Parents: []int{},
				Childs:  []int{2, 3},
			},
			{
				NodeId:  2,
				Name:    "B",
				Parents: []int{1},
				Childs:  []int{},
			},
			{
				NodeId:  3,
				Name:    "C",
				Parents: []int{1},
				Childs:  []int{},
			},
		},
	},
	{
		testCase:    2,
		testNode:    2,
		expectedVal: true,
		testTree: tree{
			{
				NodeId:  1,
				Name:    "A",
				Parents: []int{},
				Childs:  []int{2, 3},
			},
			{
				NodeId:  2,
				Name:    "B",
				Parents: []int{1},
				Childs:  []int{},
			},
			{
				NodeId:  3,
				Name:    "C",
				Parents: []int{1},
				Childs:  []int{},
			},
		},
	},
	{
		testCase:    3,
		testNode:    4,
		expectedVal: true,
		testTree: tree{
			{
				NodeId:  1,
				Name:    "A",
				Parents: []int{4},
				Childs:  []int{2, 3},
			},
			{
				NodeId:  2,
				Name:    "B",
				Parents: []int{1},
				Childs:  []int{},
			},
			{
				NodeId:  3,
				Name:    "C",
				Parents: []int{1},
				Childs:  []int{},
			},
			{
				NodeId:  4,
				Name:    "D",
				Parents: []int{},
				Childs:  []int{1},
			},
		},
	},
	{
		testCase:    4,
		testNode:    4,
		expectedVal: false,
		testTree: tree{
			{
				NodeId:  1,
				Name:    "A",
				Parents: []int{},
				Childs:  []int{2, 3},
			},
			{
				NodeId:  2,
				Name:    "B",
				Parents: []int{1},
				Childs:  []int{},
			},
			{
				NodeId:  3,
				Name:    "C",
				Parents: []int{1},
				Childs:  []int{},
			},
		},
	},
	{
		testCase:    5,
		testNode:    9,
		expectedVal: false,
		testTree: tree{
			{
				NodeId:  1,
				Name:    "A",
				Parents: []int{},
				Childs:  []int{2, 3},
			},
			{
				NodeId:  2,
				Name:    "B",
				Parents: []int{1},
				Childs:  []int{},
			},
			{
				NodeId:  3,
				Name:    "C",
				Parents: []int{1},
				Childs:  []int{},
			},
		},
	},
}

func TestValidNode(t *testing.T) {
	for _, tt := range testcases {
		actualVal := utils.ValidNode(tt.testNode, tt.testTree)
		if actualVal != tt.expectedVal {
			t.Errorf("error in test case %d", tt.testCase)
		}
	}
}
