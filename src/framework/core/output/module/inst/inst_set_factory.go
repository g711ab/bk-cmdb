/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package inst

import (
	"configcenter/src/framework/common"
	"configcenter/src/framework/core/output/module/model"
)

func createSet(target model.Model) (Inst, error) {
	return &inst{target: target}, nil
}

// findSetsLikeName find all insts by inst name
func findSetsLikeName(target model.Model, businessName string) (Iterator, error) {
	// TODO:按照名字读取特定模型的实例集合，实例名字要模糊匹配
	return &iterator{}, nil
}

// findSetsByCondition find all insts by condition
func findSetsByCondition(target model.Model, condition common.Condition) (Iterator, error) {
	// TODO:按照条件读取所有实例
	return &iterator{}, nil
}