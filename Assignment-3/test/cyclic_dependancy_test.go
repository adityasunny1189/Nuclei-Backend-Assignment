package test

import (
	"nuclei-assignment-3/utils"
	"testing"
)

var testcasesCD = []struct {
	testCase       int
	testParentNode int
	testChildNode  int
	expectedVal    bool
	testTree       tree
}{
	{
		testCase:       1,
		testParentNode: 1,
		testChildNode:  3,
		expectedVal:    false,
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
		testCase:       2,
		testParentNode: 2,
		testChildNode:  1,
		expectedVal:    true,
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
		testCase:       3,
		testParentNode: 1,
		testChildNode:  4,
		expectedVal:    true,
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
		testCase:       4,
		testParentNode: 4,
		testChildNode:  1,
		expectedVal:    true,
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
				Childs:  []int{4},
			},
			{
				NodeId:  4,
				Name:    "D",
				Parents: []int{3},
				Childs:  []int{},
			},
		},
	},
	{
		testCase:       5,
		testParentNode: 5,
		testChildNode:  1,
		expectedVal:    true,
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
				Childs:  []int{4},
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
				Parents: []int{2},
				Childs:  []int{5},
			},
			{
				NodeId:  5,
				Name:    "E",
				Parents: []int{4},
				Childs:  []int{},
			},
		},
	},
}

func TestCyclicDependancy(t *testing.T) {
	for _, tt := range testcasesCD {
		actualVal := utils.CheckCyclicDependancy(tt.testParentNode, tt.testChildNode, tt.testTree)
		if actualVal != tt.expectedVal {
			t.Errorf("error in test case %d", tt.testCase)
		}
	}
}
