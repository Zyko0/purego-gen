//go:build windows || linux

package examples

import (
	puregogen "github.com/Zyko0/purego-gen"
	purego "github.com/ebitengine/purego"
	"runtime"
	"strings"
	"unsafe"
)

var (
	// Library handles
	_hnd_cl uintptr
	// Symbols
	// cl
	_addr_clStr                              uintptr
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
	var path string

	// cl
	switch runtime.GOOS {
	case "windows":
		path = "opencl.dll"
	case "linux":
		path = "opencl.so"
	default:
		panic("os not supported: " + runtime.GOOS)
	}
	_hnd_cl, err = puregogen.OpenLibrary(path)
	if err != nil {
		panic("cannot puregogen.OpenLibrary: " + path)
	}
	// Symbols cl
	_addr_clStr, err = puregogen.OpenSymbol(_hnd_cl, "clStr")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: clStr")
	}
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

	clStr = func(s string) string {
		if !strings.HasSuffix(s, "\x00") {
			s += "\x00"
		}
		_s := uintptr(unsafe.Pointer(puregogen.BytePtrFromString(s)))
		defer runtime.KeepAlive(_s)
		_r0, _, _ := purego.SyscallN(_addr_clStr, _s)
		__r0 := "" + puregogen.BytePtrToString(*(**byte)(unsafe.Pointer(&_r0)))
		return __r0
	}
	clGetPlatformIDs = func(numEntries uint32, platforms []Platform, numPlatforms *uint32) clStatus {
		_numEntries := uintptr(numEntries)
		_platforms := uintptr(unsafe.Pointer(&platforms[0]))
		_numPlatforms := uintptr(unsafe.Pointer(numPlatforms))
		_r0, _, _ := purego.SyscallN(_addr_clGetPlatformIDs, _numEntries, _platforms, _numPlatforms)
		__r0 := clStatus(_r0)
		return __r0
	}
	clGetPlatformInfo = func(platform Platform, platformInfo platformInfo, paramValueSize clSize, paramValue []byte, paramValueSizeRet *clSize) clStatus {
		_platform := uintptr(platform)
		_platformInfo := uintptr(platformInfo)
		_paramValueSize := uintptr(paramValueSize)
		_paramValue := uintptr(unsafe.Pointer(&paramValue[0]))
		_paramValueSizeRet := uintptr(unsafe.Pointer(paramValueSizeRet))
		_r0, _, _ := purego.SyscallN(_addr_clGetPlatformInfo, _platform, _platformInfo, _paramValueSize, _paramValue, _paramValueSizeRet)
		__r0 := clStatus(_r0)
		return __r0
	}
	clGetDeviceIDs = func(platform Platform, deviceType DeviceType, numEntries uint32, devices []Device, numDevices *uint32) clStatus {
		_platform := uintptr(platform)
		_deviceType := uintptr(deviceType)
		_numEntries := uintptr(numEntries)
		_devices := uintptr(unsafe.Pointer(&devices[0]))
		_numDevices := uintptr(unsafe.Pointer(numDevices))
		_r0, _, _ := purego.SyscallN(_addr_clGetDeviceIDs, _platform, _deviceType, _numEntries, _devices, _numDevices)
		__r0 := clStatus(_r0)
		return __r0
	}
	clGetDeviceInfo = func(device Device, deviceInfo deviceInfo, paramValueSize clSize, paramValue []byte, paramValueSizeRet *clSize) clStatus {
		_device := uintptr(device)
		_deviceInfo := uintptr(deviceInfo)
		_paramValueSize := uintptr(paramValueSize)
		_paramValue := uintptr(unsafe.Pointer(&paramValue[0]))
		_paramValueSizeRet := uintptr(unsafe.Pointer(paramValueSizeRet))
		_r0, _, _ := purego.SyscallN(_addr_clGetDeviceInfo, _device, _deviceInfo, _paramValueSize, _paramValue, _paramValueSizeRet)
		__r0 := clStatus(_r0)
		return __r0
	}
	createContext = func(properties unsafe.Pointer, numDevices uint32, devices []Device, pfnNotify *createContextNotifyFunc, userData []byte, errCodeRet *clStatus) Context {
		_properties := uintptr(properties)
		_numDevices := uintptr(numDevices)
		_devices := uintptr(unsafe.Pointer(&devices[0]))
		_pfnNotify := uintptr(unsafe.Pointer(pfnNotify))
		_userData := uintptr(unsafe.Pointer(&userData[0]))
		_errCodeRet := uintptr(unsafe.Pointer(errCodeRet))
		_r0, _, _ := purego.SyscallN(_addr_clCreateContext, _properties, _numDevices, _devices, _pfnNotify, _userData, _errCodeRet)
		__r0 := Context(_r0)
		return __r0
	}
	clReleaseContext = func(ctx Context) clStatus {
		_ctx := uintptr(ctx)
		_r0, _, _ := purego.SyscallN(_addr_clReleaseContext, _ctx)
		__r0 := clStatus(_r0)
		return __r0
	}
	clCreateProgramWithSource = func(ctx Context, count clSize, strings []string, lengths []clSize, errCodeRet *clStatus) Program {
		_ctx := uintptr(ctx)
		_count := uintptr(count)
		_strings := uintptr(unsafe.Pointer(&strings[0]))
		_lengths := uintptr(unsafe.Pointer(&lengths[0]))
		_errCodeRet := uintptr(unsafe.Pointer(errCodeRet))
		_r0, _, _ := purego.SyscallN(_addr_clCreateProgramWithSource, _ctx, _count, _strings, _lengths, _errCodeRet)
		__r0 := Program(_r0)
		return __r0
	}
	clCreateBuffer = func(ctx Context, memFlags MemFlag, size clSize, hostPtr unsafe.Pointer, errCodeRet *clStatus) Buffer {
		_ctx := uintptr(ctx)
		_memFlags := uintptr(memFlags)
		_size := uintptr(size)
		_hostPtr := uintptr(hostPtr)
		_errCodeRet := uintptr(unsafe.Pointer(errCodeRet))
		_r0, _, _ := purego.SyscallN(_addr_clCreateBuffer, _ctx, _memFlags, _size, _hostPtr, _errCodeRet)
		__r0 := Buffer(_r0)
		return __r0
	}
	clCreateImage2D = func(ctx Context, memFlags MemFlag, imageFormat *ImageFormat, imageWidth clSize, imageHeight clSize, imageRowPitch clSize, hostPtr unsafe.Pointer, errCodeRet *clStatus) Buffer {
		_ctx := uintptr(ctx)
		_memFlags := uintptr(memFlags)
		_imageFormat := uintptr(unsafe.Pointer(imageFormat))
		_imageWidth := uintptr(imageWidth)
		_imageHeight := uintptr(imageHeight)
		_imageRowPitch := uintptr(imageRowPitch)
		_hostPtr := uintptr(hostPtr)
		_errCodeRet := uintptr(unsafe.Pointer(errCodeRet))
		_r0, _, _ := purego.SyscallN(_addr_clCreateImage2D, _ctx, _memFlags, _imageFormat, _imageWidth, _imageHeight, _imageRowPitch, _hostPtr, _errCodeRet)
		__r0 := Buffer(_r0)
		return __r0
	}
	clCreateCommandQueue = func(context Context, device Device, properties CommandQueueProperty, errCodeRet *clStatus) CommandQueue {
		_context := uintptr(context)
		_device := uintptr(device)
		_properties := uintptr(properties)
		_errCodeRet := uintptr(unsafe.Pointer(errCodeRet))
		_r0, _, _ := purego.SyscallN(_addr_clCreateCommandQueue, _context, _device, _properties, _errCodeRet)
		__r0 := CommandQueue(_r0)
		return __r0
	}
	clCreateCommandQueueWithProperties = func(context Context, device Device, properties CommandQueueProperty, errCodeRet *clStatus) CommandQueue {
		_context := uintptr(context)
		_device := uintptr(device)
		_properties := uintptr(properties)
		_errCodeRet := uintptr(unsafe.Pointer(errCodeRet))
		_r0, _, _ := purego.SyscallN(_addr_clCreateCommandQueueWithProperties, _context, _device, _properties, _errCodeRet)
		__r0 := CommandQueue(_r0)
		return __r0
	}
	clEnqueueBarrier = func(queue CommandQueue) clStatus {
		_queue := uintptr(queue)
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueBarrier, _queue)
		__r0 := clStatus(_r0)
		return __r0
	}
	clEnqueueNDRangeKernel = func(queue CommandQueue, kernel Kernel, workDim uint, globalWorkOffset []clSize, globalWorkSize []clSize, localWorkSize []clSize, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		_queue := uintptr(queue)
		_kernel := uintptr(kernel)
		_workDim := uintptr(workDim)
		_globalWorkOffset := uintptr(unsafe.Pointer(&globalWorkOffset[0]))
		_globalWorkSize := uintptr(unsafe.Pointer(&globalWorkSize[0]))
		_localWorkSize := uintptr(unsafe.Pointer(&localWorkSize[0]))
		_numEventsWaitList := uintptr(numEventsWaitList)
		_eventWaitList := uintptr(unsafe.Pointer(&eventWaitList[0]))
		_event := uintptr(unsafe.Pointer(event))
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueNDRangeKernel, _queue, _kernel, _workDim, _globalWorkOffset, _globalWorkSize, _localWorkSize, _numEventsWaitList, _eventWaitList, _event)
		__r0 := clStatus(_r0)
		return __r0
	}
	clEnqueueReadBuffer = func(queue CommandQueue, buffer Buffer, blockingRead bool, offset clSize, cb clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		_queue := uintptr(queue)
		_buffer := uintptr(buffer)
		_blockingRead := uintptr(0)
		if blockingRead {
			_blockingRead = 1
		}
		_offset := uintptr(offset)
		_cb := uintptr(cb)
		_ptr := uintptr(ptr)
		_numEventsWaitList := uintptr(numEventsWaitList)
		_eventWaitList := uintptr(unsafe.Pointer(&eventWaitList[0]))
		_event := uintptr(unsafe.Pointer(event))
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueReadBuffer, _queue, _buffer, _blockingRead, _offset, _cb, _ptr, _numEventsWaitList, _eventWaitList, _event)
		__r0 := clStatus(_r0)
		return __r0
	}
	clEnqueueReadImage = func(queue CommandQueue, image Buffer, blockingRead bool, origin [3]clSize, region [3]clSize, row_pitch clSize, slice_pitch clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		_queue := uintptr(queue)
		_image := uintptr(image)
		_blockingRead := uintptr(0)
		if blockingRead {
			_blockingRead = 1
		}
		_origin := uintptr(unsafe.Pointer(&origin[0]))
		_region := uintptr(unsafe.Pointer(&region[0]))
		_row_pitch := uintptr(row_pitch)
		_slice_pitch := uintptr(slice_pitch)
		_ptr := uintptr(ptr)
		_numEventsWaitList := uintptr(numEventsWaitList)
		_eventWaitList := uintptr(unsafe.Pointer(&eventWaitList[0]))
		_event := uintptr(unsafe.Pointer(event))
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueReadImage, _queue, _image, _blockingRead, _origin, _region, _row_pitch, _slice_pitch, _ptr, _numEventsWaitList, _eventWaitList, _event)
		__r0 := clStatus(_r0)
		return __r0
	}
	clEnqueueWriteBuffer = func(queue CommandQueue, buffer Buffer, blockingWrite bool, offset clSize, cb clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		_queue := uintptr(queue)
		_buffer := uintptr(buffer)
		_blockingWrite := uintptr(0)
		if blockingWrite {
			_blockingWrite = 1
		}
		_offset := uintptr(offset)
		_cb := uintptr(cb)
		_ptr := uintptr(ptr)
		_numEventsWaitList := uintptr(numEventsWaitList)
		_eventWaitList := uintptr(unsafe.Pointer(&eventWaitList[0]))
		_event := uintptr(unsafe.Pointer(event))
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueWriteBuffer, _queue, _buffer, _blockingWrite, _offset, _cb, _ptr, _numEventsWaitList, _eventWaitList, _event)
		__r0 := clStatus(_r0)
		return __r0
	}
	clEnqueueMapBuffer = func(queue CommandQueue, buffer Buffer, blockingMap bool, mapFlags MapFlag, offset clSize, size clSize, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *clStatus) uintptr {
		_queue := uintptr(queue)
		_buffer := uintptr(buffer)
		_blockingMap := uintptr(0)
		if blockingMap {
			_blockingMap = 1
		}
		_mapFlags := uintptr(mapFlags)
		_offset := uintptr(offset)
		_size := uintptr(size)
		_numEventsWaitList := uintptr(numEventsWaitList)
		_eventWaitList := uintptr(unsafe.Pointer(&eventWaitList[0]))
		_event := uintptr(unsafe.Pointer(event))
		_errCodeRet := uintptr(unsafe.Pointer(errCodeRet))
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueMapBuffer, _queue, _buffer, _blockingMap, _mapFlags, _offset, _size, _numEventsWaitList, _eventWaitList, _event, _errCodeRet)
		return _r0
	}
	clEnqueueUnmapMemObject = func(queue CommandQueue, buffer Buffer, mappedPtr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		_queue := uintptr(queue)
		_buffer := uintptr(buffer)
		_mappedPtr := uintptr(mappedPtr)
		_numEventsWaitList := uintptr(numEventsWaitList)
		_eventWaitList := uintptr(unsafe.Pointer(&eventWaitList[0]))
		_event := uintptr(unsafe.Pointer(event))
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueUnmapMemObject, _queue, _buffer, _mappedPtr, _numEventsWaitList, _eventWaitList, _event)
		__r0 := clStatus(_r0)
		return __r0
	}
	clEnqueueMapImage = func(queue CommandQueue, image Buffer, blockingMap bool, mapFlags MapFlag, origin [3]clSize, region [3]clSize, imageRowPitch *clSize, imageSlicePitch *clSize, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *clStatus) uintptr {
		_queue := uintptr(queue)
		_image := uintptr(image)
		_blockingMap := uintptr(0)
		if blockingMap {
			_blockingMap = 1
		}
		_mapFlags := uintptr(mapFlags)
		_origin := uintptr(unsafe.Pointer(&origin[0]))
		_region := uintptr(unsafe.Pointer(&region[0]))
		_imageRowPitch := uintptr(unsafe.Pointer(imageRowPitch))
		_imageSlicePitch := uintptr(unsafe.Pointer(imageSlicePitch))
		_numEventsWaitList := uintptr(numEventsWaitList)
		_eventWaitList := uintptr(unsafe.Pointer(&eventWaitList[0]))
		_event := uintptr(unsafe.Pointer(event))
		_errCodeRet := uintptr(unsafe.Pointer(errCodeRet))
		_r0, _, _ := purego.SyscallN(_addr_clEnqueueMapImage, _queue, _image, _blockingMap, _mapFlags, _origin, _region, _imageRowPitch, _imageSlicePitch, _numEventsWaitList, _eventWaitList, _event, _errCodeRet)
		return _r0
	}
	clFinishCommandQueue = func(queue CommandQueue) clStatus {
		_queue := uintptr(queue)
		_r0, _, _ := purego.SyscallN(_addr_clFinishCommandQueue, _queue)
		__r0 := clStatus(_r0)
		return __r0
	}
	clFlushCommandQueue = func(queue CommandQueue) clStatus {
		_queue := uintptr(queue)
		_r0, _, _ := purego.SyscallN(_addr_clFlushCommandQueue, _queue)
		__r0 := clStatus(_r0)
		return __r0
	}
	clReleaseCommandQueue = func(queue CommandQueue) clStatus {
		_queue := uintptr(queue)
		_r0, _, _ := purego.SyscallN(_addr_clReleaseCommandQueue, _queue)
		__r0 := clStatus(_r0)
		return __r0
	}
	clBuildProgram = func(program Program, numDevices uint32, devices []Device, options string, pfnNotify *buildProgramNotifyFunc, userData []byte) clStatus {
		_program := uintptr(program)
		_numDevices := uintptr(numDevices)
		_devices := uintptr(unsafe.Pointer(&devices[0]))
		if !strings.HasSuffix(options, "\x00") {
			options += "\x00"
		}
		_options := uintptr(unsafe.Pointer(puregogen.BytePtrFromString(options)))
		defer runtime.KeepAlive(_options)
		_pfnNotify := uintptr(unsafe.Pointer(pfnNotify))
		_userData := uintptr(unsafe.Pointer(&userData[0]))
		_r0, _, _ := purego.SyscallN(_addr_clBuildProgram, _program, _numDevices, _devices, _options, _pfnNotify, _userData)
		__r0 := clStatus(_r0)
		return __r0
	}
	clGetProgramBuildInfo = func(program Program, device Device, info programBuildInfo, paramSize clSize, paramValue unsafe.Pointer, paramSizeRet *clSize) clStatus {
		_program := uintptr(program)
		_device := uintptr(device)
		_info := uintptr(info)
		_paramSize := uintptr(paramSize)
		_paramValue := uintptr(paramValue)
		_paramSizeRet := uintptr(unsafe.Pointer(paramSizeRet))
		_r0, _, _ := purego.SyscallN(_addr_clGetProgramBuildInfo, _program, _device, _info, _paramSize, _paramValue, _paramSizeRet)
		__r0 := clStatus(_r0)
		return __r0
	}
	clCreateKernel = func(program Program, kernelName string, errCodeRet *clStatus) Kernel {
		_program := uintptr(program)
		if !strings.HasSuffix(kernelName, "\x00") {
			kernelName += "\x00"
		}
		_kernelName := uintptr(unsafe.Pointer(puregogen.BytePtrFromString(kernelName)))
		defer runtime.KeepAlive(_kernelName)
		_errCodeRet := uintptr(unsafe.Pointer(errCodeRet))
		_r0, _, _ := purego.SyscallN(_addr_clCreateKernel, _program, _kernelName, _errCodeRet)
		__r0 := Kernel(_r0)
		return __r0
	}
	clReleaseProgram = func(program Program) clStatus {
		_program := uintptr(program)
		_r0, _, _ := purego.SyscallN(_addr_clReleaseProgram, _program)
		__r0 := clStatus(_r0)
		return __r0
	}
	clSetKernelArg = func(kernel Kernel, argIndex uint, argSize clSize, argValue unsafe.Pointer) clStatus {
		_kernel := uintptr(kernel)
		_argIndex := uintptr(argIndex)
		_argSize := uintptr(argSize)
		_argValue := uintptr(argValue)
		_r0, _, _ := purego.SyscallN(_addr_clSetKernelArg, _kernel, _argIndex, _argSize, _argValue)
		__r0 := clStatus(_r0)
		return __r0
	}
	clReleaseKernel = func(kernel Kernel) clStatus {
		_kernel := uintptr(kernel)
		_r0, _, _ := purego.SyscallN(_addr_clReleaseKernel, _kernel)
		__r0 := clStatus(_r0)
		return __r0
	}
	clGetMemObjectInfo = func(buffer Buffer, memInfo memInfo, paramValueSize clSize, paramValue unsafe.Pointer, paramValueSizeRet *clSize) clStatus {
		_buffer := uintptr(buffer)
		_memInfo := uintptr(memInfo)
		_paramValueSize := uintptr(paramValueSize)
		_paramValue := uintptr(paramValue)
		_paramValueSizeRet := uintptr(unsafe.Pointer(paramValueSizeRet))
		_r0, _, _ := purego.SyscallN(_addr_clGetMemObjectInfo, _buffer, _memInfo, _paramValueSize, _paramValue, _paramValueSizeRet)
		__r0 := clStatus(_r0)
		return __r0
	}
	clReleaseMemObject = func(buffer Buffer) clStatus {
		_buffer := uintptr(buffer)
		_r0, _, _ := purego.SyscallN(_addr_clReleaseMemObject, _buffer)
		__r0 := clStatus(_r0)
		return __r0
	}
}
