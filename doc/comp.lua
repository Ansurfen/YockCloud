-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class complib
---@field dial fun(configuration: string): comp

---@class comp
---@field exec fun(cmd: string)
---@field query fun(cmd: string)
---@field close fun(): err
---@field ping fun(): err
