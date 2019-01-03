// Copyright 2019 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// Package gwheel provides Timing Wheel for interval jobs running and management/时间轮.
// 高效的时间轮任务执行管理，用于管理异步的间隔运行任务，或者异步只运行一次的任务(默认最小时间粒度为秒)。
// 与其他定时任务管理模块的区别是，时间轮模块只管理间隔执行任务，并且更注重执行效率(纳秒级别)。
package gwheel

import "time"

const (
    STATUS_READY            = 0
    STATUS_RUNNING          = 1
    STATUS_CLOSED           = -1
    gPANIC_EXIT             = "exit"
    gDEFAULT_SLOT_NUMBER    = 10
    gDEFAULT_WHEEL_INTERVAL = 10*time.Millisecond
)

var (
    // 默认的wheel管理对象
    defaultWheel = NewDefault()
)

// 添加执行方法，可以给定名字，以便于后续执行删除
func Add(interval time.Duration, job JobFunc) (*Entry, error) {
    return defaultWheel.Add(interval, job)
}

// 添加单例运行循环任务
func AddSingleton(interval time.Duration, job JobFunc) (*Entry, error) {
    return defaultWheel.AddSingleton(interval, job)
}

// 添加只运行一次的循环任务
func AddOnce(interval time.Duration, job JobFunc) (*Entry, error) {
    return defaultWheel.AddOnce(interval, job)
}

// 添加运行指定次数的循环任务
func AddTimes(interval time.Duration, times int, job JobFunc) (*Entry, error) {
    return defaultWheel.AddTimes(interval, times, job)
}

// 延迟添加循环任务，delay参数单位为秒
func DelayAdd(delay time.Duration, interval time.Duration, job JobFunc) {
    defaultWheel.DelayAdd(delay, interval, job)
}

// 延迟添加单例循环任务，delay参数单位为秒
func DelayAddSingleton(delay time.Duration, interval time.Duration, job JobFunc) {
    defaultWheel.DelayAddSingleton(delay, interval, job)
}

// 延迟添加只运行一次的循环任务，delay参数单位为秒
func DelayAddOnce(delay time.Duration, interval time.Duration, job JobFunc) {
    defaultWheel.DelayAddOnce(delay, interval, job)
}

// 延迟添加运行指定次数的循环任务，delay参数单位为秒
func DelayAddTimes(delay time.Duration, interval time.Duration, times int, job JobFunc) {
    defaultWheel.DelayAddTimes(delay, interval, times, job)
}

// 在Job方法中调用，停止当前运行的任务
func Exit() {
    panic(gPANIC_EXIT)
}