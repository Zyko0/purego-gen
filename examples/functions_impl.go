//go:build windows || linux

package examples

import (
	puregogen "github.com/Zyko0/purego-gen"
	purego "github.com/ebitengine/purego"
	"runtime"
	"unsafe"
)

// File generated by github.com/Zyko0/purego-gen. DO NOT EDIT.

var (
	// Library handles
	_hnd_cl uintptr
	// Symbols
	// cl
	_addr_clGetPlatformIDs                   uintptr
	_addr_clGetPlatformInfo                  uintptr
	_addr_clGetDeviceIDs                     uintptr
	_addr_clGetDeviceInfo                    uintptr
	_addr_clCreateContext                    uintptr
	_addr_clReleaseContext                   uintptr
	_addr_clCreateProgramWithSource          uintptr
	_addr_clCreateBuffer                     uintptr
	_addr_clCreateImage2D                    uintptr
	_addr_clCreateCommandQueue               uintptr
	_addr_clCreateCommandQueueWithProperties uintptr
	_addr_clEnqueueBarrier                   uintptr
	_addr_clEnqueueNDRangeKernel             uintptr
	_addr_clEnqueueReadBuffer                uintptr
	_addr_clEnqueueReadImage                 uintptr
	_addr_clEnqueueWriteBuffer               uintptr
	_addr_clEnqueueMapBuffer                 uintptr
	_addr_clEnqueueUnmapMemObject            uintptr
	_addr_clEnqueueMapImage                  uintptr
	_addr_clFinishCommandQueue               uintptr
	_addr_clFlushCommandQueue                uintptr
	_addr_clReleaseCommandQueue              uintptr
	_addr_clBuildProgram                     uintptr
	_addr_clGetProgramBuildInfo              uintptr
	_addr_clCreateKernel                     uintptr
	_addr_clReleaseProgram                   uintptr
	_addr_clSetKernelArg                     uintptr
	_addr_clReleaseKernel                    uintptr
	_addr_clGetMemObjectInfo                 uintptr
	_addr_clReleaseMemObject                 uintptr
)

func init() {
	var err error

	// Symbols cl
	_addr_clGetPlatformIDs, err = puregogen.OpenSymbol(_hnd_cl, "clGetPlatformIDs")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clGetPlatformIDs")
	}
	_addr_clGetPlatformInfo, err = puregogen.OpenSymbol(_hnd_cl, "clGetPlatformInfo")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clGetPlatformInfo")
	}
	_addr_clGetDeviceIDs, err = puregogen.OpenSymbol(_hnd_cl, "clGetDeviceIDs")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clGetDeviceIDs")
	}
	_addr_clGetDeviceInfo, err = puregogen.OpenSymbol(_hnd_cl, "clGetDeviceInfo")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clGetDeviceInfo")
	}
	_addr_clCreateContext, err = puregogen.OpenSymbol(_hnd_cl, "clCreateContext")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clCreateContext")
	}
	_addr_clReleaseContext, err = puregogen.OpenSymbol(_hnd_cl, "clReleaseContext")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clReleaseContext")
	}
	_addr_clCreateProgramWithSource, err = puregogen.OpenSymbol(_hnd_cl, "clCreateProgramWithSource")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clCreateProgramWithSource")
	}
	_addr_clCreateBuffer, err = puregogen.OpenSymbol(_hnd_cl, "clCreateBuffer")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clCreateBuffer")
	}
	_addr_clCreateImage2D, err = puregogen.OpenSymbol(_hnd_cl, "clCreateImage2D")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clCreateImage2D")
	}
	_addr_clCreateCommandQueue, err = puregogen.OpenSymbol(_hnd_cl, "clCreateCommandQueue")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clCreateCommandQueue")
	}
	_addr_clCreateCommandQueueWithProperties, err = puregogen.OpenSymbol(_hnd_cl, "clCreateCommandQueueWithProperties")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clCreateCommandQueueWithProperties")
	}
	_addr_clEnqueueBarrier, err = puregogen.OpenSymbol(_hnd_cl, "clEnqueueBarrier")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clEnqueueBarrier")
	}
	_addr_clEnqueueNDRangeKernel, err = puregogen.OpenSymbol(_hnd_cl, "clEnqueueNDRangeKernel")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clEnqueueNDRangeKernel")
	}
	_addr_clEnqueueReadBuffer, err = puregogen.OpenSymbol(_hnd_cl, "clEnqueueReadBuffer")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clEnqueueReadBuffer")
	}
	_addr_clEnqueueReadImage, err = puregogen.OpenSymbol(_hnd_cl, "clEnqueueReadImage")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clEnqueueReadImage")
	}
	_addr_clEnqueueWriteBuffer, err = puregogen.OpenSymbol(_hnd_cl, "clEnqueueWriteBuffer")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clEnqueueWriteBuffer")
	}
	_addr_clEnqueueMapBuffer, err = puregogen.OpenSymbol(_hnd_cl, "clEnqueueMapBuffer")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clEnqueueMapBuffer")
	}
	_addr_clEnqueueUnmapMemObject, err = puregogen.OpenSymbol(_hnd_cl, "clEnqueueUnmapMemObject")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clEnqueueUnmapMemObject")
	}
	_addr_clEnqueueMapImage, err = puregogen.OpenSymbol(_hnd_cl, "clEnqueueMapImage")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clEnqueueMapImage")
	}
	_addr_clFinishCommandQueue, err = puregogen.OpenSymbol(_hnd_cl, "clFinishCommandQueue")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clFinishCommandQueue")
	}
	_addr_clFlushCommandQueue, err = puregogen.OpenSymbol(_hnd_cl, "clFlushCommandQueue")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clFlushCommandQueue")
	}
	_addr_clReleaseCommandQueue, err = puregogen.OpenSymbol(_hnd_cl, "clReleaseCommandQueue")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clReleaseCommandQueue")
	}
	_addr_clBuildProgram, err = puregogen.OpenSymbol(_hnd_cl, "clBuildProgram")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clBuildProgram")
	}
	_addr_clGetProgramBuildInfo, err = puregogen.OpenSymbol(_hnd_cl, "clGetProgramBuildInfo")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clGetProgramBuildInfo")
	}
	_addr_clCreateKernel, err = puregogen.OpenSymbol(_hnd_cl, "clCreateKernel")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clCreateKernel")
	}
	_addr_clReleaseProgram, err = puregogen.OpenSymbol(_hnd_cl, "clReleaseProgram")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clReleaseProgram")
	}
	_addr_clSetKernelArg, err = puregogen.OpenSymbol(_hnd_cl, "clSetKernelArg")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clSetKernelArg")
	}
	_addr_clReleaseKernel, err = puregogen.OpenSymbol(_hnd_cl, "clReleaseKernel")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clReleaseKernel")
	}
	_addr_clGetMemObjectInfo, err = puregogen.OpenSymbol(_hnd_cl, "clGetMemObjectInfo")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clGetMemObjectInfo")
	}
	_addr_clReleaseMemObject, err = puregogen.OpenSymbol(_hnd_cl, "clReleaseMemObject")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clReleaseMemObject")
	}

	clGetPlatformIDs = func(numEntries uint32, platforms []Platform, numPlatforms *uint32) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clGetPlatformIDs, uintptr(numEntries), uintptr(unsafe.Pointer(unsafe.SliceData(platforms))), uintptr(unsafe.Pointer(numPlatforms)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(platforms)
		runtime.KeepAlive(numPlatforms)
		return __r0
	}
	clGetPlatformInfo = func(platform Platform, platformInfo platformInfo, paramValueSize clSize, paramValue []byte, paramValueSizeRet *clSize) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clGetPlatformInfo, uintptr(platform), uintptr(platformInfo), uintptr(paramValueSize), uintptr(unsafe.Pointer(unsafe.SliceData(paramValue))), uintptr(unsafe.Pointer(paramValueSizeRet)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(paramValue)
		runtime.KeepAlive(paramValueSizeRet)
		return __r0
	}
	clGetDeviceIDs = func(platform Platform, deviceType DeviceType, numEntries uint32, devices []Device, numDevices *uint32) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clGetDeviceIDs, uintptr(platform), uintptr(deviceType), uintptr(numEntries), uintptr(unsafe.Pointer(unsafe.SliceData(devices))), uintptr(unsafe.Pointer(numDevices)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(devices)
		runtime.KeepAlive(numDevices)
		return __r0
	}
	clGetDeviceInfo = func(device Device, deviceInfo deviceInfo, paramValueSize clSize, paramValue []byte, paramValueSizeRet *clSize) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clGetDeviceInfo, uintptr(device), uintptr(deviceInfo), uintptr(paramValueSize), uintptr(unsafe.Pointer(unsafe.SliceData(paramValue))), uintptr(unsafe.Pointer(paramValueSizeRet)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(paramValue)
		runtime.KeepAlive(paramValueSizeRet)
		return __r0
	}
	createContext = func(properties unsafe.Pointer, numDevices uint32, devices []Device, pfnNotify *createContextNotifyFunc, userData []byte, errCodeRet *clStatus) Context {
		_r0, _, _ := purego.SyscallN(_addr_clCreateContext, uintptr(properties), uintptr(numDevices), uintptr(unsafe.Pointer(unsafe.SliceData(devices))), uintptr(unsafe.Pointer(pfnNotify)), uintptr(unsafe.Pointer(unsafe.SliceData(userData))), uintptr(unsafe.Pointer(errCodeRet)))
		__r0 := Context(_r0)
		runtime.KeepAlive(devices)
		runtime.KeepAlive(pfnNotify)
		runtime.KeepAlive(userData)
		runtime.KeepAlive(errCodeRet)
		return __r0
	}
	clReleaseContext = func(ctx Context) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clReleaseContext, uintptr(ctx))
		__r0 := clStatus(_r0)
		return __r0
	}
	clCreateProgramWithSource = func(ctx Context, count clSize, strings []string, lengths []clSize, errCodeRet *clStatus) Program {
		_r0, _, _ := purego.SyscallN(_addr_clCreateProgramWithSource, uintptr(ctx), uintptr(count), uintptr(unsafe.Pointer(unsafe.SliceData(strings))), uintptr(unsafe.Pointer(unsafe.SliceData(lengths))), uintptr(unsafe.Pointer(errCodeRet)))
		__r0 := Program(_r0)
		runtime.KeepAlive(strings)
		runtime.KeepAlive(lengths)
		runtime.KeepAlive(errCodeRet)
		return __r0
	}
	clCreateBuffer = func(ctx Context, memFlags MemFlag, size clSize, hostPtr unsafe.Pointer, errCodeRet *clStatus) Buffer {
		_r0, _, _ := purego.SyscallN(_addr_clCreateBuffer, uintptr(ctx), uintptr(memFlags), uintptr(size), uintptr(hostPtr), uintptr(unsafe.Pointer(errCodeRet)))
		__r0 := Buffer(_r0)
		runtime.KeepAlive(errCodeRet)
		return __r0
	}
	clCreateImage2D = func(ctx Context, memFlags MemFlag, imageFormat *ImageFormat, imageWidth clSize, imageHeight clSize, imageRowPitch clSize, hostPtr unsafe.Pointer, errCodeRet *clStatus) Buffer {
		_r0, _, _ := purego.SyscallN(_addr_clCreateImage2D, uintptr(ctx), uintptr(memFlags), uintptr(unsafe.Pointer(imageFormat)), uintptr(imageWidth), uintptr(imageHeight), uintptr(imageRowPitch), uintptr(hostPtr), uintptr(unsafe.Pointer(errCodeRet)))
		__r0 := Buffer(_r0)
		runtime.KeepAlive(imageFormat)
		runtime.KeepAlive(errCodeRet)
		return __r0
	}
	clCreateCommandQueue = func(context Context, device Device, properties CommandQueueProperty, errCodeRet *clStatus) CommandQueue {
		_r0, _, _ := purego.SyscallN(_addr_clCreateCommandQueue, uintptr(context), uintptr(device), uintptr(properties), uintptr(unsafe.Pointer(errCodeRet)))
		__r0 := CommandQueue(_r0)
		runtime.KeepAlive(errCodeRet)
		return __r0
	}
	clCreateCommandQueueWithProperties = func(context Context, device Device, properties CommandQueueProperty, errCodeRet *clStatus) CommandQueue {
		_r0, _, _ := purego.SyscallN(_addr_clCreateCommandQueueWithProperties, uintptr(context), uintptr(device), uintptr(properties), uintptr(unsafe.Pointer(errCodeRet)))
		__r0 := CommandQueue(_r0)
		runtime.KeepAlive(errCodeRet)
		return __r0
	}
	clEnqueueBarrier = func(queue CommandQueue) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueBarrier, uintptr(queue))
		__r0 := clStatus(_r0)
		return __r0
	}
	clEnqueueNDRangeKernel = func(queue CommandQueue, kernel Kernel, workDim uint, globalWorkOffset []clSize, globalWorkSize []clSize, localWorkSize []clSize, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueNDRangeKernel, uintptr(queue), uintptr(kernel), uintptr(workDim), uintptr(unsafe.Pointer(unsafe.SliceData(globalWorkOffset))), uintptr(unsafe.Pointer(unsafe.SliceData(globalWorkSize))), uintptr(unsafe.Pointer(unsafe.SliceData(localWorkSize))), uintptr(numEventsWaitList), uintptr(unsafe.Pointer(unsafe.SliceData(eventWaitList))), uintptr(unsafe.Pointer(event)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(globalWorkOffset)
		runtime.KeepAlive(globalWorkSize)
		runtime.KeepAlive(localWorkSize)
		runtime.KeepAlive(eventWaitList)
		runtime.KeepAlive(event)
		return __r0
	}
	clEnqueueReadBuffer = func(queue CommandQueue, buffer Buffer, blockingRead bool, offset clSize, cb clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueReadBuffer, uintptr(queue), uintptr(buffer), puregogen.BoolToUintptr(blockingRead), uintptr(offset), uintptr(cb), uintptr(ptr), uintptr(numEventsWaitList), uintptr(unsafe.Pointer(unsafe.SliceData(eventWaitList))), uintptr(unsafe.Pointer(event)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(eventWaitList)
		runtime.KeepAlive(event)
		return __r0
	}
	clEnqueueReadImage = func(queue CommandQueue, image Buffer, blockingRead bool, origin [3]clSize, region [3]clSize, row_pitch clSize, slice_pitch clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueReadImage, uintptr(queue), uintptr(image), puregogen.BoolToUintptr(blockingRead), uintptr(unsafe.Pointer(unsafe.SliceData(origin[:]))), uintptr(unsafe.Pointer(unsafe.SliceData(region[:]))), uintptr(row_pitch), uintptr(slice_pitch), uintptr(ptr), uintptr(numEventsWaitList), uintptr(unsafe.Pointer(unsafe.SliceData(eventWaitList))), uintptr(unsafe.Pointer(event)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(origin)
		runtime.KeepAlive(region)
		runtime.KeepAlive(eventWaitList)
		runtime.KeepAlive(event)
		return __r0
	}
	clEnqueueWriteBuffer = func(queue CommandQueue, buffer Buffer, blockingWrite bool, offset clSize, cb clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueWriteBuffer, uintptr(queue), uintptr(buffer), puregogen.BoolToUintptr(blockingWrite), uintptr(offset), uintptr(cb), uintptr(ptr), uintptr(numEventsWaitList), uintptr(unsafe.Pointer(unsafe.SliceData(eventWaitList))), uintptr(unsafe.Pointer(event)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(eventWaitList)
		runtime.KeepAlive(event)
		return __r0
	}
	clEnqueueMapBuffer = func(queue CommandQueue, buffer Buffer, blockingMap bool, mapFlags MapFlag, offset clSize, size clSize, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *clStatus) uintptr {
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueMapBuffer, uintptr(queue), uintptr(buffer), puregogen.BoolToUintptr(blockingMap), uintptr(mapFlags), uintptr(offset), uintptr(size), uintptr(numEventsWaitList), uintptr(unsafe.Pointer(unsafe.SliceData(eventWaitList))), uintptr(unsafe.Pointer(event)), uintptr(unsafe.Pointer(errCodeRet)))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(eventWaitList)
		runtime.KeepAlive(event)
		runtime.KeepAlive(errCodeRet)
		return __r0
	}
	clEnqueueUnmapMemObject = func(queue CommandQueue, buffer Buffer, mappedPtr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueUnmapMemObject, uintptr(queue), uintptr(buffer), uintptr(mappedPtr), uintptr(numEventsWaitList), uintptr(unsafe.Pointer(unsafe.SliceData(eventWaitList))), uintptr(unsafe.Pointer(event)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(eventWaitList)
		runtime.KeepAlive(event)
		return __r0
	}
	clEnqueueMapImage = func(queue CommandQueue, image Buffer, blockingMap bool, mapFlags MapFlag, origin [3]clSize, region [3]clSize, imageRowPitch *clSize, imageSlicePitch *clSize, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *clStatus) uintptr {
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueMapImage, uintptr(queue), uintptr(image), puregogen.BoolToUintptr(blockingMap), uintptr(mapFlags), uintptr(unsafe.Pointer(unsafe.SliceData(origin[:]))), uintptr(unsafe.Pointer(unsafe.SliceData(region[:]))), uintptr(unsafe.Pointer(imageRowPitch)), uintptr(unsafe.Pointer(imageSlicePitch)), uintptr(numEventsWaitList), uintptr(unsafe.Pointer(unsafe.SliceData(eventWaitList))), uintptr(unsafe.Pointer(event)), uintptr(unsafe.Pointer(errCodeRet)))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(origin)
		runtime.KeepAlive(region)
		runtime.KeepAlive(imageRowPitch)
		runtime.KeepAlive(imageSlicePitch)
		runtime.KeepAlive(eventWaitList)
		runtime.KeepAlive(event)
		runtime.KeepAlive(errCodeRet)
		return __r0
	}
	clFinishCommandQueue = func(queue CommandQueue) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clFinishCommandQueue, uintptr(queue))
		__r0 := clStatus(_r0)
		return __r0
	}
	clFlushCommandQueue = func(queue CommandQueue) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clFlushCommandQueue, uintptr(queue))
		__r0 := clStatus(_r0)
		return __r0
	}
	clReleaseCommandQueue = func(queue CommandQueue) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clReleaseCommandQueue, uintptr(queue))
		__r0 := clStatus(_r0)
		return __r0
	}
	clBuildProgram = func(program Program, numDevices uint32, devices []Device, options string, pfnNotify *buildProgramNotifyFunc, userData []byte) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clBuildProgram, uintptr(program), uintptr(numDevices), uintptr(unsafe.Pointer(unsafe.SliceData(devices))), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(options))), uintptr(unsafe.Pointer(pfnNotify)), uintptr(unsafe.Pointer(unsafe.SliceData(userData))))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(devices)
		runtime.KeepAlive(options)
		runtime.KeepAlive(pfnNotify)
		runtime.KeepAlive(userData)
		return __r0
	}
	clGetProgramBuildInfo = func(program Program, device Device, info programBuildInfo, paramSize clSize, paramValue unsafe.Pointer, paramSizeRet *clSize) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clGetProgramBuildInfo, uintptr(program), uintptr(device), uintptr(info), uintptr(paramSize), uintptr(paramValue), uintptr(unsafe.Pointer(paramSizeRet)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(paramSizeRet)
		return __r0
	}
	clCreateKernel = func(program Program, kernelName string, errCodeRet *clStatus) Kernel {
		_r0, _, _ := purego.SyscallN(_addr_clCreateKernel, uintptr(program), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(kernelName))), uintptr(unsafe.Pointer(errCodeRet)))
		__r0 := Kernel(_r0)
		runtime.KeepAlive(kernelName)
		runtime.KeepAlive(errCodeRet)
		return __r0
	}
	clReleaseProgram = func(program Program) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clReleaseProgram, uintptr(program))
		__r0 := clStatus(_r0)
		return __r0
	}
	clSetKernelArg = func(kernel Kernel, argIndex uint, argSize clSize, argValue unsafe.Pointer) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clSetKernelArg, uintptr(kernel), uintptr(argIndex), uintptr(argSize), uintptr(argValue))
		__r0 := clStatus(_r0)
		return __r0
	}
	clReleaseKernel = func(kernel Kernel) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clReleaseKernel, uintptr(kernel))
		__r0 := clStatus(_r0)
		return __r0
	}
	clGetMemObjectInfo = func(buffer Buffer, memInfo memInfo, paramValueSize clSize, paramValue unsafe.Pointer, paramValueSizeRet *clSize) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clGetMemObjectInfo, uintptr(buffer), uintptr(memInfo), uintptr(paramValueSize), uintptr(paramValue), uintptr(unsafe.Pointer(paramValueSizeRet)))
		__r0 := clStatus(_r0)
		runtime.KeepAlive(paramValueSizeRet)
		return __r0
	}
	clReleaseMemObject = func(buffer Buffer) clStatus {
		_r0, _, _ := purego.SyscallN(_addr_clReleaseMemObject, uintptr(buffer))
		__r0 := clStatus(_r0)
		return __r0
	}
}
