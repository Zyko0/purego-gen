package examples

import (
	"unsafe"
)

type (
	Platform                uint32
	platformInfo            uint
	Device                  uint32
	DeviceType              uint32
	deviceInfo              uint
	Context                 uint32
	clStatus                uint32
	clSize                  uint32
	CommandQueueProperty    uint32
	CommandQueue            uint32
	Kernel                  uint32
	Program                 uint32
	MemFlag                 uint32
	ImageFormat             uint32
	Buffer                  uint32
	programBuildInfo        uint32
	MapFlag                 uint32
	buildProgramNotifyFunc  uintptr
	createContextNotifyFunc uintptr
	Event                   uint32
	memInfo                 uint32
)

//go:generate purego-gen --input ./examples/functions.go

var (
	//puregogen:library path:windows=opencl.dll path:linux=opencl.so alias=cl
	clGetPlatformIDs  func(numEntries uint32, platforms []Platform, numPlatforms *uint32) clStatus
	clGetPlatformInfo func(platform Platform, platformInfo platformInfo, paramValueSize clSize, paramValue []byte, paramValueSizeRet *clSize) clStatus
	// Device
	clGetDeviceIDs  func(platform Platform, deviceType DeviceType, numEntries uint32, devices []Device, numDevices *uint32) clStatus
	clGetDeviceInfo func(device Device, deviceInfo deviceInfo, paramValueSize clSize, paramValue []byte, paramValueSizeRet *clSize) clStatus
	// Context
	//puregogen:function symbol=clCreateContext
	createContext             func(properties unsafe.Pointer, numDevices uint32, devices []Device, pfnNotify *createContextNotifyFunc, userData []byte, errCodeRet *clStatus) Context
	clReleaseContext          func(ctx Context) clStatus
	clCreateProgramWithSource func(ctx Context, count clSize, strings []string, lengths []clSize, errCodeRet *clStatus) Program
	clCreateBuffer            func(ctx Context, memFlags MemFlag, size clSize, hostPtr unsafe.Pointer, errCodeRet *clStatus) Buffer
	clCreateImage2D           func(ctx Context, memFlags MemFlag, imageFormat *ImageFormat, imageWidth, imageHeight, imageRowPitch clSize, hostPtr unsafe.Pointer, errCodeRet *clStatus) Buffer
	// Command queue
	clCreateCommandQueue               func(context Context, device Device, properties CommandQueueProperty, errCodeRet *clStatus) CommandQueue
	clCreateCommandQueueWithProperties func(context Context, device Device, properties CommandQueueProperty, errCodeRet *clStatus) CommandQueue
	clEnqueueBarrier                   func(queue CommandQueue) clStatus
	clEnqueueNDRangeKernel             func(queue CommandQueue, kernel Kernel, workDim uint, globalWorkOffset, globalWorkSize, localWorkSize []clSize, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	clEnqueueReadBuffer                func(queue CommandQueue, buffer Buffer, blockingRead bool, offset, cb clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	clEnqueueReadImage                 func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]clSize, row_pitch, slice_pitch clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	clEnqueueWriteBuffer               func(queue CommandQueue, buffer Buffer, blockingWrite bool, offset, cb clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	clEnqueueMapBuffer                 func(queue CommandQueue, buffer Buffer, blockingMap bool, mapFlags MapFlag, offset, size clSize, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *clStatus) uintptr
	clEnqueueUnmapMemObject            func(queue CommandQueue, buffer Buffer, mappedPtr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	clEnqueueMapImage                  func(queue CommandQueue, image Buffer, blockingMap bool, mapFlags MapFlag, origin, region [3]clSize, imageRowPitch, imageSlicePitch *clSize, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *clStatus) uintptr
	clFinishCommandQueue               func(queue CommandQueue) clStatus
	clFlushCommandQueue                func(queue CommandQueue) clStatus
	clReleaseCommandQueue              func(queue CommandQueue) clStatus
	// Program
	clBuildProgram        func(program Program, numDevices uint32, devices []Device, options string, pfnNotify *buildProgramNotifyFunc, userData []byte) clStatus
	clGetProgramBuildInfo func(program Program, device Device, info programBuildInfo, paramSize clSize, paramValue unsafe.Pointer, paramSizeRet *clSize) clStatus
	clCreateKernel        func(program Program, kernelName string, errCodeRet *clStatus) Kernel
	clReleaseProgram      func(program Program) clStatus
	// Kernel
	clSetKernelArg  func(kernel Kernel, argIndex uint, argSize clSize, argValue unsafe.Pointer) clStatus
	clReleaseKernel func(kernel Kernel) clStatus
	// Buffer
	clGetMemObjectInfo func(buffer Buffer, memInfo memInfo, paramValueSize clSize, paramValue unsafe.Pointer, paramValueSizeRet *clSize) clStatus
	clReleaseMemObject func(buffer Buffer) clStatus
)
