// Copyright 2020 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

// Body Event-specific information.
type Body struct {
	// If 'allThreadsContinued' is true, a debug adapter can announce that all threads have continued.
	AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`

	// The thread which was continued.
	ThreadId float64 `json:"threadId,omitempty"`
}

// AttachRequestArguments Arguments for 'attach' request. Additional attributes are implementation specific.
type AttachRequestArguments struct {
	// Optional data from the previous, restarted session.
	// The data is sent as the 'restart' attribute of the 'terminated' event.
	// The client should leave the data intact.
	Restart interface{} `json:"__restart,omitempty"`
}

// AttachRequest Attach request; value of command field is 'attach'.
// The attach request is sent from the client to the debug adapter to attach to a debuggee that is already running. Since attaching is debugger/runtime specific, the arguments for this request are not part of this specification.
type AttachRequest struct {
	// Object containing arguments for the command.
	Arguments *AttachRequestArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// AttachResponse Response to 'attach' request. This is just an acknowledgement, so no body field is required.
type AttachResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Breakpoint Information about a Breakpoint created in setBreakpoints or setFunctionBreakpoints.
type Breakpoint struct {
	// An optional start column of the actual range covered by the breakpoint.
	Column float64 `json:"column,omitempty"`

	// An optional end column of the actual range covered by the breakpoint. If no end line is given, then the end column is assumed to be in the start line.
	EndColumn float64 `json:"endColumn,omitempty"`

	// An optional end line of the actual range covered by the breakpoint.
	EndLine float64 `json:"endLine,omitempty"`

	// An optional identifier for the breakpoint. It is needed if breakpoint events are used to update or remove breakpoints.
	Id float64 `json:"id,omitempty"`

	// The start line of the actual range covered by the breakpoint.
	Line float64 `json:"line,omitempty"`

	// An optional message about the state of the breakpoint. This is shown to the user and can be used to explain why a breakpoint could not be verified.
	Message string `json:"message,omitempty"`

	// The source where the breakpoint is located.
	Source *Source `json:"source,omitempty"`

	// If true breakpoint could be set (but not necessarily at the desired location).
	Verified bool `json:"verified,omitempty"`
}

// BreakpointEvent Event message for 'breakpoint' event type.
// The event indicates that some information about a breakpoint has changed.
type BreakpointEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// BreakpointLocation Properties of a breakpoint location returned from the 'breakpointLocations' request.
type BreakpointLocation struct {
	// Optional start column of breakpoint location.
	Column float64 `json:"column,omitempty"`

	// Optional end column of breakpoint location if the location covers a range.
	EndColumn float64 `json:"endColumn,omitempty"`

	// Optional end line of breakpoint location if the location covers a range.
	EndLine float64 `json:"endLine,omitempty"`

	// Start line of breakpoint location.
	Line float64 `json:"line,omitempty"`
}

// BreakpointLocationsArguments Arguments for 'breakpointLocations' request.
type BreakpointLocationsArguments struct {
	// Optional start column of range to search possible breakpoint locations in. If no start column is given, the first column in the start line is assumed.
	Column float64 `json:"column,omitempty"`

	// Optional end column of range to search possible breakpoint locations in. If no end column is given, then it is assumed to be in the last column of the end line.
	EndColumn float64 `json:"endColumn,omitempty"`

	// Optional end line of range to search possible breakpoint locations in. If no end line is given, then the end line is assumed to be the start line.
	EndLine float64 `json:"endLine,omitempty"`

	// Start line of range to search possible breakpoint locations in. If only the line is specified, the request returns all possible locations in that line.
	Line float64 `json:"line,omitempty"`

	// The source location of the breakpoints; either 'source.path' or 'source.reference' must be specified.
	Source *Source `json:"source,omitempty"`
}

// BreakpointLocationsRequest BreakpointLocations request; value of command field is 'breakpointLocations'.
// The 'breakpointLocations' request returns all possible locations for source breakpoints in a given range.
type BreakpointLocationsRequest struct {
	// Object containing arguments for the command.
	Arguments *BreakpointLocationsArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// BreakpointLocationsResponse Response to 'breakpointLocations' request.
// Contains possible locations for source breakpoints.
type BreakpointLocationsResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// CancelArguments Arguments for 'cancel' request.
type CancelArguments struct {
	// The ID (attribute 'seq') of the request to cancel.
	RequestId float64 `json:"requestId,omitempty"`
}

// CancelRequest Cancel request; value of command field is 'cancel'.
// The 'cancel' request is used by the frontend to indicate that it is no longer interested in the result produced by a specific request issued earlier.
// This request has a hint characteristic: a debug adapter can only be expected to make a 'best effort' in honouring this request but there are no guarantees.
// The 'cancel' request may return an error if it could not cancel an operation but a frontend should refrain from presenting this error to end users.
// A frontend client should only call this request if the capability 'supportsCancelRequest' is true.
// The request that got canceled still needs to send a response back.
// This can either be a normal result ('success' attribute true) or an error response ('success' attribute false and the 'message' set to 'cancelled').
// Returning partial results from a cancelled request is possible but please note that a frontend client has no generic way for detecting that a response is partial or not.
type CancelRequest struct {
	// Object containing arguments for the command.
	Arguments *CancelArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// CancelResponse Response to 'cancel' request. This is just an acknowledgement, so no body field is required.
type CancelResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Capabilities Information about the capabilities of a debug adapter.
type Capabilities struct {
	// The set of additional module information exposed by the debug adapter.
	AdditionalModuleColumns []*ColumnDescriptor `json:"additionalModuleColumns,omitempty"`

	// The set of characters that should trigger completion in a REPL. If not specified, the UI should assume the '.' character.
	CompletionTriggerCharacters []string `json:"completionTriggerCharacters,omitempty"`

	// Available filters or options for the setExceptionBreakpoints request.
	ExceptionBreakpointFilters []*ExceptionBreakpointsFilter `json:"exceptionBreakpointFilters,omitempty"`

	// The debug adapter supports the 'terminateDebuggee' attribute on the 'disconnect' request.
	SupportTerminateDebuggee bool `json:"supportTerminateDebuggee,omitempty"`

	// Checksum algorithms supported by the debug adapter.
	SupportedChecksumAlgorithms []string `json:"supportedChecksumAlgorithms,omitempty"`

	// The debug adapter supports the 'breakpointLocations' request.
	SupportsBreakpointLocationsRequest bool `json:"supportsBreakpointLocationsRequest,omitempty"`

	// The debug adapter supports the 'cancel' request.
	SupportsCancelRequest bool `json:"supportsCancelRequest,omitempty"`

	// The debug adapter supports the 'completions' request.
	SupportsCompletionsRequest bool `json:"supportsCompletionsRequest,omitempty"`

	// The debug adapter supports conditional breakpoints.
	SupportsConditionalBreakpoints bool `json:"supportsConditionalBreakpoints,omitempty"`

	// The debug adapter supports the 'configurationDone' request.
	SupportsConfigurationDoneRequest bool `json:"supportsConfigurationDoneRequest,omitempty"`

	// The debug adapter supports data breakpoints.
	SupportsDataBreakpoints bool `json:"supportsDataBreakpoints,omitempty"`

	// The debug adapter supports the delayed loading of parts of the stack, which requires that both the 'startFrame' and 'levels' arguments and the 'totalFrames' result of the 'StackTrace' request are supported.
	SupportsDelayedStackTraceLoading bool `json:"supportsDelayedStackTraceLoading,omitempty"`

	// The debug adapter supports the 'disassemble' request.
	SupportsDisassembleRequest bool `json:"supportsDisassembleRequest,omitempty"`

	// The debug adapter supports a (side effect free) evaluate request for data hovers.
	SupportsEvaluateForHovers bool `json:"supportsEvaluateForHovers,omitempty"`

	// The debug adapter supports the 'exceptionInfo' request.
	SupportsExceptionInfoRequest bool `json:"supportsExceptionInfoRequest,omitempty"`

	// The debug adapter supports 'exceptionOptions' on the setExceptionBreakpoints request.
	SupportsExceptionOptions bool `json:"supportsExceptionOptions,omitempty"`

	// The debug adapter supports function breakpoints.
	SupportsFunctionBreakpoints bool `json:"supportsFunctionBreakpoints,omitempty"`

	// The debug adapter supports the 'gotoTargets' request.
	SupportsGotoTargetsRequest bool `json:"supportsGotoTargetsRequest,omitempty"`

	// The debug adapter supports breakpoints that break execution after a specified number of hits.
	SupportsHitConditionalBreakpoints bool `json:"supportsHitConditionalBreakpoints,omitempty"`

	// The debug adapter supports the 'loadedSources' request.
	SupportsLoadedSourcesRequest bool `json:"supportsLoadedSourcesRequest,omitempty"`

	// The debug adapter supports logpoints by interpreting the 'logMessage' attribute of the SourceBreakpoint.
	SupportsLogPoints bool `json:"supportsLogPoints,omitempty"`

	// The debug adapter supports the 'modules' request.
	SupportsModulesRequest bool `json:"supportsModulesRequest,omitempty"`

	// The debug adapter supports the 'readMemory' request.
	SupportsReadMemoryRequest bool `json:"supportsReadMemoryRequest,omitempty"`

	// The debug adapter supports restarting a frame.
	SupportsRestartFrame bool `json:"supportsRestartFrame,omitempty"`

	// The debug adapter supports the 'restart' request. In this case a client should not implement 'restart' by terminating and relaunching the adapter but by calling the RestartRequest.
	SupportsRestartRequest bool `json:"supportsRestartRequest,omitempty"`

	// The debug adapter supports the 'setExpression' request.
	SupportsSetExpression bool `json:"supportsSetExpression,omitempty"`

	// The debug adapter supports setting a variable to a value.
	SupportsSetVariable bool `json:"supportsSetVariable,omitempty"`

	// The debug adapter supports stepping back via the 'stepBack' and 'reverseContinue' requests.
	SupportsStepBack bool `json:"supportsStepBack,omitempty"`

	// The debug adapter supports the 'stepInTargets' request.
	SupportsStepInTargetsRequest bool `json:"supportsStepInTargetsRequest,omitempty"`

	// The debug adapter supports the 'terminate' request.
	SupportsTerminateRequest bool `json:"supportsTerminateRequest,omitempty"`

	// The debug adapter supports the 'terminateThreads' request.
	SupportsTerminateThreadsRequest bool `json:"supportsTerminateThreadsRequest,omitempty"`

	// The debug adapter supports a 'format' attribute on the stackTraceRequest, variablesRequest, and evaluateRequest.
	SupportsValueFormattingOptions bool `json:"supportsValueFormattingOptions,omitempty"`
}

// CapabilitiesEvent Event message for 'capabilities' event type.
// The event indicates that one or more capabilities have changed.
// Since the capabilities are dependent on the frontend and its UI, it might not be possible to change that at random times (or too late).
// Consequently this event has a hint characteristic: a frontend can only be expected to make a 'best effort' in honouring individual capabilities but there are no guarantees.
// Only changed capabilities need to be included, all other capabilities keep their values.
type CapabilitiesEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Checksum The checksum of an item calculated by the specified algorithm.
type Checksum struct {
	// The algorithm used to calculate this checksum.
	Algorithm string `json:"algorithm,omitempty"`

	// Value of the checksum.
	Checksum string `json:"checksum,omitempty"`
}

// ColumnDescriptor A ColumnDescriptor specifies what module attribute to show in a column of the ModulesView, how to format it, and what the column's label should be.
// It is only used if the underlying UI actually supports this level of customization.
type ColumnDescriptor struct {
	// Name of the attribute rendered in this column.
	AttributeName string `json:"attributeName,omitempty"`

	// Format to use for the rendered values in this column. TBD how the format strings looks like.
	Format string `json:"format,omitempty"`

	// Header UI label of column.
	Label string `json:"label,omitempty"`

	// Datatype of values in this column.  Defaults to 'string' if not specified.
	Type string `json:"type,omitempty"`

	// Width of this column in characters (hint only).
	Width float64 `json:"width,omitempty"`
}

// CompletionItem CompletionItems are the suggestions returned from the CompletionsRequest.
type CompletionItem struct {
	// The label of this completion item. By default this is also the text that is inserted when selecting this completion.
	Label string `json:"label,omitempty"`

	// This value determines how many characters are overwritten by the completion text.
	// If missing the value 0 is assumed which results in the completion text being inserted.
	Length float64 `json:"length,omitempty"`

	// A string that should be used when comparing this item with other items. When `falsy` the label is used.
	SortText string `json:"sortText,omitempty"`

	// This value determines the location (in the CompletionsRequest's 'text' attribute) where the completion text is added.
	// If missing the text is added at the location specified by the CompletionsRequest's 'column' attribute.
	Start float64 `json:"start,omitempty"`

	// If text is not falsy then it is inserted instead of the label.
	Text string `json:"text,omitempty"`

	// The item's type. Typically the client uses this information to render the item in the UI with an icon.
	Type string `json:"type,omitempty"`
}

// CompletionsArguments Arguments for 'completions' request.
type CompletionsArguments struct {
	// The character position for which to determine the completion proposals.
	Column float64 `json:"column,omitempty"`

	// Returns completions in the scope of this stack frame. If not specified, the completions are returned for the global scope.
	FrameId float64 `json:"frameId,omitempty"`

	// An optional line for which to determine the completion proposals. If missing the first line of the text is assumed.
	Line float64 `json:"line,omitempty"`

	// One or more source lines. Typically this is the text a user has typed into the debug console before he asked for completion.
	Text string `json:"text,omitempty"`
}

// CompletionsRequest Completions request; value of command field is 'completions'.
// Returns a list of possible completions for a given caret position and text.
// The CompletionsRequest may only be called if the 'supportsCompletionsRequest' capability exists and is true.
type CompletionsRequest struct {
	// Object containing arguments for the command.
	Arguments *CompletionsArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// CompletionsResponse Response to 'completions' request.
type CompletionsResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ConfigurationDoneArguments Arguments for 'configurationDone' request.
type ConfigurationDoneArguments struct{}

// ConfigurationDoneRequest ConfigurationDone request; value of command field is 'configurationDone'.
// The client of the debug protocol must send this request at the end of the sequence of configuration requests (which was started by the 'initialized' event).
type ConfigurationDoneRequest struct {
	// Object containing arguments for the command.
	Arguments *ConfigurationDoneArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ConfigurationDoneResponse Response to 'configurationDone' request. This is just an acknowledgement, so no body field is required.
type ConfigurationDoneResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ContinueArguments Arguments for 'continue' request.
type ContinueArguments struct {
	// Continue execution for the specified thread (if possible). If the backend cannot continue on a single thread but will continue on all threads, it should set the 'allThreadsContinued' attribute in the response to true.
	ThreadId float64 `json:"threadId,omitempty"`
}

// ContinueRequest Continue request; value of command field is 'continue'.
// The request starts the debuggee to run again.
type ContinueRequest struct {
	// Object containing arguments for the command.
	Arguments *ContinueArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ContinueResponse Response to 'continue' request.
type ContinueResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ContinuedEvent Event message for 'continued' event type.
// The event indicates that the execution of the debuggee has continued.
// Please note: a debug adapter is not expected to send this event in response to a request that implies that execution continues, e.g. 'launch' or 'continue'.
// It is only necessary to send a 'continued' event if there was no previous request that implied this.
type ContinuedEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// DataBreakpoint Properties of a data breakpoint passed to the setDataBreakpoints request.
type DataBreakpoint struct {
	// The access type of the data.
	AccessType string `json:"accessType,omitempty"`

	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

	// An id representing the data. This id is returned from the dataBreakpointInfo request.
	DataId string `json:"dataId,omitempty"`

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`
}

// DataBreakpointInfoArguments Arguments for 'dataBreakpointInfo' request.
type DataBreakpointInfoArguments struct {
	// The name of the Variable's child to obtain data breakpoint information for. If variableReference isn’t provided, this can be an expression.
	Name string `json:"name,omitempty"`

	// Reference to the Variable container if the data breakpoint is requested for a child of the container.
	VariablesReference float64 `json:"variablesReference,omitempty"`
}

// DataBreakpointInfoRequest DataBreakpointInfo request; value of command field is 'dataBreakpointInfo'.
// Obtains information on a possible data breakpoint that could be set on an expression or variable.
type DataBreakpointInfoRequest struct {
	// Object containing arguments for the command.
	Arguments *DataBreakpointInfoArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// DataBreakpointInfoResponse Response to 'dataBreakpointInfo' request.
type DataBreakpointInfoResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// DisassembleArguments Arguments for 'disassemble' request.
type DisassembleArguments struct {
	// Number of instructions to disassemble starting at the specified location and offset. An adapter must return exactly this number of instructions - any unavailable instructions should be replaced with an implementation-defined 'invalid instruction' value.
	InstructionCount float64 `json:"instructionCount,omitempty"`

	// Optional offset (in instructions) to be applied after the byte offset (if any) before disassembling. Can be negative.
	InstructionOffset float64 `json:"instructionOffset,omitempty"`

	// Memory reference to the base location containing the instructions to disassemble.
	MemoryReference string `json:"memoryReference,omitempty"`

	// Optional offset (in bytes) to be applied to the reference location before disassembling. Can be negative.
	Offset float64 `json:"offset,omitempty"`

	// If true, the adapter should attempt to resolve memory addresses and other values to symbolic names.
	ResolveSymbols bool `json:"resolveSymbols,omitempty"`
}

// DisassembleRequest Disassemble request; value of command field is 'disassemble'.
// Disassembles code stored at the provided location.
type DisassembleRequest struct {
	// Object containing arguments for the command.
	Arguments *DisassembleArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// DisassembleResponse Response to 'disassemble' request.
type DisassembleResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// DisassembledInstruction Represents a single disassembled instruction.
type DisassembledInstruction struct {
	// The address of the instruction. Treated as a hex value if prefixed with '0x', or as a decimal value otherwise.
	Address string `json:"address,omitempty"`

	// The column within the line that corresponds to this instruction, if any.
	Column float64 `json:"column,omitempty"`

	// The end column of the range that corresponds to this instruction, if any.
	EndColumn float64 `json:"endColumn,omitempty"`

	// The end line of the range that corresponds to this instruction, if any.
	EndLine float64 `json:"endLine,omitempty"`

	// Text representing the instruction and its operands, in an implementation-defined format.
	Instruction string `json:"instruction,omitempty"`

	// Optional raw bytes representing the instruction and its operands, in an implementation-defined format.
	InstructionBytes string `json:"instructionBytes,omitempty"`

	// The line within the source location that corresponds to this instruction, if any.
	Line float64 `json:"line,omitempty"`

	// Source location that corresponds to this instruction, if any. Should always be set (if available) on the first instruction returned, but can be omitted afterwards if this instruction maps to the same source file as the previous instruction.
	Location *Source `json:"location,omitempty"`

	// Name of the symbol that corresponds with the location of this instruction, if any.
	Symbol string `json:"symbol,omitempty"`
}

// DisconnectArguments Arguments for 'disconnect' request.
type DisconnectArguments struct {
	// A value of true indicates that this 'disconnect' request is part of a restart sequence.
	Restart bool `json:"restart,omitempty"`

	// Indicates whether the debuggee should be terminated when the debugger is disconnected.
	// If unspecified, the debug adapter is free to do whatever it thinks is best.
	// A client can only rely on this attribute being properly honored if a debug adapter returns true for the 'supportTerminateDebuggee' capability.
	TerminateDebuggee bool `json:"terminateDebuggee,omitempty"`
}

// DisconnectRequest Disconnect request; value of command field is 'disconnect'.
// The 'disconnect' request is sent from the client to the debug adapter in order to stop debugging. It asks the debug adapter to disconnect from the debuggee and to terminate the debug adapter. If the debuggee has been started with the 'launch' request, the 'disconnect' request terminates the debuggee. If the 'attach' request was used to connect to the debuggee, 'disconnect' does not terminate the debuggee. This behavior can be controlled with the 'terminateDebuggee' argument (if supported by the debug adapter).
type DisconnectRequest struct {
	// Object containing arguments for the command.
	Arguments *DisconnectArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// DisconnectResponse Response to 'disconnect' request. This is just an acknowledgement, so no body field is required.
type DisconnectResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ErrorResponse On error (whenever 'success' is false), the body can provide more details.
type ErrorResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// EvaluateArguments Arguments for 'evaluate' request.
type EvaluateArguments struct {
	// The context in which the evaluate request is run.
	// Values:
	// 'watch': evaluate is run in a watch.
	// 'repl': evaluate is run from REPL console.
	// 'hover': evaluate is run from a data hover.
	// etc.
	Context string `json:"context,omitempty"`

	// The expression to evaluate.
	Expression string `json:"expression,omitempty"`

	// Specifies details on how to format the Evaluate result.
	Format *ValueFormat `json:"format,omitempty"`

	// Evaluate the expression in the scope of this stack frame. If not specified, the expression is evaluated in the global scope.
	FrameId float64 `json:"frameId,omitempty"`
}

// EvaluateRequest Evaluate request; value of command field is 'evaluate'.
// Evaluates the given expression in the context of the top most stack frame.
// The expression has access to any variables and arguments that are in scope.
type EvaluateRequest struct {
	// Object containing arguments for the command.
	Arguments *EvaluateArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// EvaluateResponse Response to 'evaluate' request.
type EvaluateResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Event A debug adapter initiated event.
type Event struct {
	// Event-specific information.
	Body interface{} `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ExceptionBreakpointsFilter An ExceptionBreakpointsFilter is shown in the UI as an option for configuring how exceptions are dealt with.
type ExceptionBreakpointsFilter struct {
	// Initial value of the filter. If not specified a value 'false' is assumed.
	Default bool `json:"default,omitempty"`

	// The internal ID of the filter. This value is passed to the setExceptionBreakpoints request.
	Filter string `json:"filter,omitempty"`

	// The name of the filter. This will be shown in the UI.
	Label string `json:"label,omitempty"`
}

// ExceptionDetails Detailed information about an exception that has occurred.
type ExceptionDetails struct {
	// Optional expression that can be evaluated in the current scope to obtain the exception object.
	EvaluateName string `json:"evaluateName,omitempty"`

	// Fully-qualified type name of the exception object.
	FullTypeName string `json:"fullTypeName,omitempty"`

	// Details of the exception contained by this exception, if any.
	InnerException []*ExceptionDetails `json:"innerException,omitempty"`

	// Message contained in the exception.
	Message string `json:"message,omitempty"`

	// Stack trace at the time the exception was thrown.
	StackTrace string `json:"stackTrace,omitempty"`

	// Short type name of the exception object.
	TypeName string `json:"typeName,omitempty"`
}

// ExceptionInfoArguments Arguments for 'exceptionInfo' request.
type ExceptionInfoArguments struct {
	// Thread for which exception information should be retrieved.
	ThreadId float64 `json:"threadId,omitempty"`
}

// ExceptionInfoRequest ExceptionInfo request; value of command field is 'exceptionInfo'.
// Retrieves the details of the exception that caused this event to be raised.
type ExceptionInfoRequest struct {
	// Object containing arguments for the command.
	Arguments *ExceptionInfoArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ExceptionInfoResponse Response to 'exceptionInfo' request.
type ExceptionInfoResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ExceptionOptions An ExceptionOptions assigns configuration options to a set of exceptions.
type ExceptionOptions struct {
	// Condition when a thrown exception should result in a break.
	BreakMode string `json:"breakMode,omitempty"`

	// A path that selects a single or multiple exceptions in a tree. If 'path' is missing, the whole tree is selected. By convention the first segment of the path is a category that is used to group exceptions in the UI.
	Path []*ExceptionPathSegment `json:"path,omitempty"`
}

// ExceptionPathSegment An ExceptionPathSegment represents a segment in a path that is used to match leafs or nodes in a tree of exceptions. If a segment consists of more than one name, it matches the names provided if 'negate' is false or missing or it matches anything except the names provided if 'negate' is true.
type ExceptionPathSegment struct {
	// Depending on the value of 'negate' the names that should match or not match.
	Names []string `json:"names,omitempty"`

	// If false or missing this segment matches the names provided, otherwise it matches anything except the names provided.
	Negate bool `json:"negate,omitempty"`
}

// ExitedEvent Event message for 'exited' event type.
// The event indicates that the debuggee has exited and returns its exit code.
type ExitedEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// FunctionBreakpoint Properties of a breakpoint passed to the setFunctionBreakpoints request.
type FunctionBreakpoint struct {
	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`

	// The name of the function.
	Name string `json:"name,omitempty"`
}

// GotoArguments Arguments for 'goto' request.
type GotoArguments struct {
	// The location where the debuggee will continue to run.
	TargetId float64 `json:"targetId,omitempty"`

	// Set the goto target for this thread.
	ThreadId float64 `json:"threadId,omitempty"`
}

// GotoRequest Goto request; value of command field is 'goto'.
// The request sets the location where the debuggee will continue to run.
// This makes it possible to skip the execution of code or to executed code again.
// The code between the current location and the goto target is not executed but skipped.
// The debug adapter first sends the response and then a 'stopped' event with reason 'goto'.
type GotoRequest struct {
	// Object containing arguments for the command.
	Arguments *GotoArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// GotoResponse Response to 'goto' request. This is just an acknowledgement, so no body field is required.
type GotoResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// GotoTarget A GotoTarget describes a code location that can be used as a target in the 'goto' request.
// The possible goto targets can be determined via the 'gotoTargets' request.
type GotoTarget struct {
	// An optional column of the goto target.
	Column float64 `json:"column,omitempty"`

	// An optional end column of the range covered by the goto target.
	EndColumn float64 `json:"endColumn,omitempty"`

	// An optional end line of the range covered by the goto target.
	EndLine float64 `json:"endLine,omitempty"`

	// Unique identifier for a goto target. This is used in the goto request.
	Id float64 `json:"id,omitempty"`

	// Optional memory reference for the instruction pointer value represented by this target.
	InstructionPointerReference string `json:"instructionPointerReference,omitempty"`

	// The name of the goto target (shown in the UI).
	Label string `json:"label,omitempty"`

	// The line of the goto target.
	Line float64 `json:"line,omitempty"`
}

// GotoTargetsArguments Arguments for 'gotoTargets' request.
type GotoTargetsArguments struct {
	// An optional column location for which the goto targets are determined.
	Column float64 `json:"column,omitempty"`

	// The line location for which the goto targets are determined.
	Line float64 `json:"line,omitempty"`

	// The source location for which the goto targets are determined.
	Source *Source `json:"source,omitempty"`
}

// GotoTargetsRequest GotoTargets request; value of command field is 'gotoTargets'.
// This request retrieves the possible goto targets for the specified source location.
// These targets can be used in the 'goto' request.
// The GotoTargets request may only be called if the 'supportsGotoTargetsRequest' capability exists and is true.
type GotoTargetsRequest struct {
	// Object containing arguments for the command.
	Arguments *GotoTargetsArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// GotoTargetsResponse Response to 'gotoTargets' request.
type GotoTargetsResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// InitializeRequestArguments Arguments for 'initialize' request.
type InitializeRequestArguments struct {
	// The ID of the debug adapter.
	AdapterID string `json:"adapterID,omitempty"`

	// The ID of the (frontend) client using this adapter.
	ClientID string `json:"clientID,omitempty"`

	// The human readable name of the (frontend) client using this adapter.
	ClientName string `json:"clientName,omitempty"`

	// If true all column numbers are 1-based (default).
	ColumnsStartAt1 bool `json:"columnsStartAt1,omitempty"`

	// If true all line numbers are 1-based (default).
	LinesStartAt1 bool `json:"linesStartAt1,omitempty"`

	// The ISO-639 locale of the (frontend) client using this adapter, e.g. en-US or de-CH.
	Locale string `json:"locale,omitempty"`

	// Determines in what format paths are specified. The default is 'path', which is the native format.
	// Values: 'path', 'uri', etc.
	PathFormat string `json:"pathFormat,omitempty"`

	// Client supports memory references.
	SupportsMemoryReferences bool `json:"supportsMemoryReferences,omitempty"`

	// Client supports the runInTerminal request.
	SupportsRunInTerminalRequest bool `json:"supportsRunInTerminalRequest,omitempty"`

	// Client supports the paging of variables.
	SupportsVariablePaging bool `json:"supportsVariablePaging,omitempty"`

	// Client supports the optional type attribute for variables.
	SupportsVariableType bool `json:"supportsVariableType,omitempty"`
}

// InitializeRequest Initialize request; value of command field is 'initialize'.
// The 'initialize' request is sent as the first request from the client to the debug adapter in order to configure it with client capabilities and to retrieve capabilities from the debug adapter.
// Until the debug adapter has responded to with an 'initialize' response, the client must not send any additional requests or events to the debug adapter. In addition the debug adapter is not allowed to send any requests or events to the client until it has responded with an 'initialize' response.
// The 'initialize' request may only be sent once.
type InitializeRequest struct {
	// Object containing arguments for the command.
	Arguments *InitializeRequestArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// InitializeResponse Response to 'initialize' request.
type InitializeResponse struct {
	// The capabilities of this debug adapter.
	Body *Capabilities `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// InitializedEvent Event message for 'initialized' event type.
// This event indicates that the debug adapter is ready to accept configuration requests (e.g. SetBreakpointsRequest, SetExceptionBreakpointsRequest).
// A debug adapter is expected to send this event when it is ready to accept configuration requests (but not before the 'initialize' request has finished).
// The sequence of events/requests is as follows:
// - adapters sends 'initialized' event (after the 'initialize' request has returned)
// - frontend sends zero or more 'setBreakpoints' requests
// - frontend sends one 'setFunctionBreakpoints' request
// - frontend sends a 'setExceptionBreakpoints' request if one or more 'exceptionBreakpointFilters' have been defined (or if 'supportsConfigurationDoneRequest' is not defined or false)
// - frontend sends other future configuration requests
// - frontend sends one 'configurationDone' request to indicate the end of the configuration.
type InitializedEvent struct {
	// Event-specific information.
	Body interface{} `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// LaunchRequest Launch request; value of command field is 'launch'.
// The launch request is sent from the client to the debug adapter to start the debuggee with or without debugging (if 'noDebug' is true). Since launching is debugger/runtime specific, the arguments for this request are not part of this specification.
type LaunchRequest struct {
	// Object containing arguments for the command.
	Arguments *LaunchRequestArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// LaunchRequestArguments Arguments for 'launch' request. Additional attributes are implementation specific.
type LaunchRequestArguments struct {
	// If noDebug is true the launch request should launch the program without enabling debugging.
	NoDebug bool `json:"noDebug,omitempty"`

	// Optional data from the previous, restarted session.
	// The data is sent as the 'restart' attribute of the 'terminated' event.
	// The client should leave the data intact.
	Restart interface{} `json:"__restart,omitempty"`
}

// LaunchResponse Response to 'launch' request. This is just an acknowledgement, so no body field is required.
type LaunchResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// LoadedSourceEvent Event message for 'loadedSource' event type.
// The event indicates that some source has been added, changed, or removed from the set of all loaded sources.
type LoadedSourceEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// LoadedSourcesArguments Arguments for 'loadedSources' request.
type LoadedSourcesArguments struct {
}

// LoadedSourcesRequest LoadedSources request; value of command field is 'loadedSources'.
// Retrieves the set of all sources currently loaded by the debugged process.
type LoadedSourcesRequest struct {
	// Object containing arguments for the command.
	Arguments *LoadedSourcesArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// LoadedSourcesResponse Response to 'loadedSources' request.
type LoadedSourcesResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Message A structured message object. Used to return errors from requests.
type Message struct {
	// A format string for the message. Embedded variables have the form '{name}'.
	// If variable name starts with an underscore character, the variable does not contain user data (PII) and can be safely used for telemetry purposes.
	Format string `json:"format,omitempty"`

	// Unique identifier for the message.
	Id float64 `json:"id,omitempty"`

	// If true send to telemetry.
	SendTelemetry bool `json:"sendTelemetry,omitempty"`

	// If true show user.
	ShowUser bool `json:"showUser,omitempty"`

	// An optional url where additional information about this message can be found.
	Url string `json:"url,omitempty"`

	// An optional label that is presented to the user as the UI for opening the url.
	UrlLabel string `json:"urlLabel,omitempty"`

	// An object used as a dictionary for looking up the variables in the format string.
	Variables map[string]string `json:"variables,omitempty"`
}

// Module A Module object represents a row in the modules view.
// Two attributes are mandatory: an id identifies a module in the modules view and is used in a ModuleEvent for identifying a module for adding, updating or deleting.
// The name is used to minimally render the module in the UI.
//
// Additional attributes can be added to the module. They will show up in the module View if they have a corresponding ColumnDescriptor.
//
// To avoid an unnecessary proliferation of additional attributes with similar semantics but different names
// we recommend to re-use attributes from the 'recommended' list below first, and only introduce new attributes if nothing appropriate could be found.
type Module struct {
	// Address range covered by this module.
	AddressRange string `json:"addressRange,omitempty"`

	// Module created or modified.
	DateTimeStamp string `json:"dateTimeStamp,omitempty"`

	// Unique identifier for the module.
	Id interface{} `json:"id,omitempty"`

	// True if the module is optimized.
	IsOptimized bool `json:"isOptimized,omitempty"`

	// True if the module is considered 'user code' by a debugger that supports 'Just My Code'.
	IsUserCode bool `json:"isUserCode,omitempty"`

	// A name of the module.
	Name string `json:"name,omitempty"`

	// optional but recommended attributes.
	// always try to use these first before introducing additional attributes.
	//
	// Logical full path to the module. The exact definition is implementation defined, but usually this would be a full path to the on-disk file for the module.
	Path string `json:"path,omitempty"`

	// Logical full path to the symbol file. The exact definition is implementation defined.
	SymbolFilePath string `json:"symbolFilePath,omitempty"`

	// User understandable description of if symbols were found for the module (ex: 'Symbols Loaded', 'Symbols not found', etc.
	SymbolStatus string `json:"symbolStatus,omitempty"`

	// Version of Module.
	Version string `json:"version,omitempty"`
}

// ModuleEvent Event message for 'module' event type.
// The event indicates that some information about a module has changed.
type ModuleEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ModulesArguments Arguments for 'modules' request.
type ModulesArguments struct {
	// The number of modules to return. If moduleCount is not specified or 0, all modules are returned.
	ModuleCount float64 `json:"moduleCount,omitempty"`

	// The index of the first module to return; if omitted modules start at 0.
	StartModule float64 `json:"startModule,omitempty"`
}

// ModulesRequest Modules request; value of command field is 'modules'.
// Modules can be retrieved from the debug adapter with the ModulesRequest which can either return all modules or a range of modules to support paging.
type ModulesRequest struct {
	// Object containing arguments for the command.
	Arguments *ModulesArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ModulesResponse Response to 'modules' request.
type ModulesResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ModulesViewDescriptor The ModulesViewDescriptor is the container for all declarative configuration options of a ModuleView.
// For now it only specifies the columns to be shown in the modules view.
type ModulesViewDescriptor struct {
	Columns []*ColumnDescriptor `json:"columns,omitempty"`
}

// NextArguments Arguments for 'next' request.
type NextArguments struct {
	// Execute 'next' for this thread.
	ThreadId float64 `json:"threadId,omitempty"`
}

// NextRequest Next request; value of command field is 'next'.
// The request starts the debuggee to run again for one step.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'step') after the step has completed.
type NextRequest struct {
	// Object containing arguments for the command.
	Arguments *NextArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// NextResponse Response to 'next' request. This is just an acknowledgement, so no body field is required.
type NextResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// OutputEvent Event message for 'output' event type.
// The event indicates that the target has produced some output.
type OutputEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// PauseArguments Arguments for 'pause' request.
type PauseArguments struct {
	// Pause execution for this thread.
	ThreadId float64 `json:"threadId,omitempty"`
}

// PauseRequest Pause request; value of command field is 'pause'.
// The request suspends the debuggee.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'pause') after the thread has been paused successfully.
type PauseRequest struct {
	// Object containing arguments for the command.
	Arguments *PauseArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// PauseResponse Response to 'pause' request. This is just an acknowledgement, so no body field is required.
type PauseResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ProcessEvent Event message for 'process' event type.
// The event indicates that the debugger has begun debugging a new process. Either one that it has launched, or one that it has attached to.
type ProcessEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ProtocolMessage Base class of requests, responses, and events.
type ProtocolMessage struct {
	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ReadMemoryArguments Arguments for 'readMemory' request.
type ReadMemoryArguments struct {
	// Number of bytes to read at the specified location and offset.
	Count float64 `json:"count,omitempty"`

	// Memory reference to the base location from which data should be read.
	MemoryReference string `json:"memoryReference,omitempty"`

	// Optional offset (in bytes) to be applied to the reference location before reading data. Can be negative.
	Offset float64 `json:"offset,omitempty"`
}

// ReadMemoryRequest ReadMemory request; value of command field is 'readMemory'.
// Reads bytes from memory at the provided location.
type ReadMemoryRequest struct {
	// Object containing arguments for the command.
	Arguments *ReadMemoryArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ReadMemoryResponse Response to 'readMemory' request.
type ReadMemoryResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Request A client or debug adapter initiated request.
type Request struct {
	// Object containing arguments for the command.
	Arguments interface{} `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Response Response for a request.
type Response struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// RestartArguments Arguments for 'restart' request.
type RestartArguments struct {
}

// RestartFrameArguments Arguments for 'restartFrame' request.
type RestartFrameArguments struct {
	// Restart this stackframe.
	FrameId float64 `json:"frameId,omitempty"`
}

// RestartFrameRequest RestartFrame request; value of command field is 'restartFrame'.
// The request restarts execution of the specified stackframe.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'restart') after the restart has completed.
type RestartFrameRequest struct {
	// Object containing arguments for the command.
	Arguments *RestartFrameArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// RestartFrameResponse Response to 'restartFrame' request. This is just an acknowledgement, so no body field is required.
type RestartFrameResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// RestartRequest Restart request; value of command field is 'restart'.
// Restarts a debug session. If the capability 'supportsRestartRequest' is missing or has the value false,
// the client will implement 'restart' by terminating the debug adapter first and then launching it anew.
// A debug adapter can override this default behaviour by implementing a restart request
// and setting the capability 'supportsRestartRequest' to true.
type RestartRequest struct {
	// Object containing arguments for the command.
	Arguments *RestartArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// RestartResponse Response to 'restart' request. This is just an acknowledgement, so no body field is required.
type RestartResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ReverseContinueArguments Arguments for 'reverseContinue' request.
type ReverseContinueArguments struct {
	// Execute 'reverseContinue' for this thread.
	ThreadId float64 `json:"threadId,omitempty"`
}

// ReverseContinueRequest ReverseContinue request; value of command field is 'reverseContinue'.
// The request starts the debuggee to run backward. Clients should only call this request if the capability 'supportsStepBack' is true.
type ReverseContinueRequest struct {
	// Object containing arguments for the command.
	Arguments *ReverseContinueArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ReverseContinueResponse Response to 'reverseContinue' request. This is just an acknowledgement, so no body field is required.
type ReverseContinueResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// RunInTerminalRequestArguments Arguments for 'runInTerminal' request.
type RunInTerminalRequestArguments struct {
	// List of arguments. The first argument is the command to run.
	Args []string `json:"args,omitempty"`

	// Working directory of the command.
	Cwd string `json:"cwd,omitempty"`

	// Environment key-value pairs that are added to or removed from the default environment.
	Env map[string]string `json:"env,omitempty"`

	// What kind of terminal to launch.
	Kind string `json:"kind,omitempty"`

	// Optional title of the terminal.
	Title string `json:"title,omitempty"`
}

// RunInTerminalRequest RunInTerminal request; value of command field is 'runInTerminal'.
// This request is sent from the debug adapter to the client to run a command in a terminal. This is typically used to launch the debuggee in a terminal provided by the client.
type RunInTerminalRequest struct {
	// Object containing arguments for the command.
	Arguments *RunInTerminalRequestArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// RunInTerminalResponse Response to 'runInTerminal' request.
type RunInTerminalResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Scope A Scope is a named container for variables. Optionally a scope can map to a source or a range within a source.
type Scope struct {
	// Optional start column of the range covered by this scope.
	Column float64 `json:"column,omitempty"`

	// Optional end column of the range covered by this scope.
	EndColumn float64 `json:"endColumn,omitempty"`

	// Optional end line of the range covered by this scope.
	EndLine float64 `json:"endLine,omitempty"`

	// If true, the number of variables in this scope is large or expensive to retrieve.
	Expensive bool `json:"expensive,omitempty"`

	// The number of indexed variables in this scope.
	// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
	IndexedVariables float64 `json:"indexedVariables,omitempty"`

	// Optional start line of the range covered by this scope.
	Line float64 `json:"line,omitempty"`

	// Name of the scope such as 'Arguments', 'Locals', or 'Registers'. This string is shown in the UI as is and can be translated.
	Name string `json:"name,omitempty"`

	// The number of named variables in this scope.
	// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
	NamedVariables float64 `json:"namedVariables,omitempty"`

	// An optional hint for how to present this scope in the UI. If this attribute is missing, the scope is shown with a generic UI.
	// Values:
	// 'arguments': Scope contains method arguments.
	// 'locals': Scope contains local variables.
	// 'registers': Scope contains registers. Only a single 'registers' scope should be returned from a 'scopes' request.
	// etc.
	PresentationHint string `json:"presentationHint,omitempty"`

	// Optional source for this scope.
	Source *Source `json:"source,omitempty"`

	// The variables of this scope can be retrieved by passing the value of variablesReference to the VariablesRequest.
	VariablesReference float64 `json:"variablesReference,omitempty"`
}

// ScopesArguments Arguments for 'scopes' request.
type ScopesArguments struct {
	// Retrieve the scopes for this stackframe.
	FrameId float64 `json:"frameId,omitempty"`
}

// ScopesRequest Scopes request; value of command field is 'scopes'.
// The request returns the variable scopes for a given stackframe ID.
type ScopesRequest struct {
	// Object containing arguments for the command.
	Arguments *ScopesArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ScopesResponse Response to 'scopes' request.
type ScopesResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetBreakpointsArguments Arguments for 'setBreakpoints' request.
type SetBreakpointsArguments struct {
	// The code locations of the breakpoints.
	Breakpoints []*SourceBreakpoint `json:"breakpoints,omitempty"`

	// Deprecated: The code locations of the breakpoints.
	Lines []float64 `json:"lines,omitempty"`

	// The source location of the breakpoints; either 'source.path' or 'source.reference' must be specified.
	Source *Source `json:"source,omitempty"`

	// A value of true indicates that the underlying source has been modified which results in new breakpoint locations.
	SourceModified bool `json:"sourceModified,omitempty"`
}

// SetBreakpointsRequest SetBreakpoints request; value of command field is 'setBreakpoints'.
// Sets multiple breakpoints for a single source and clears all previous breakpoints in that source.
// To clear all breakpoint for a source, specify an empty array.
// When a breakpoint is hit, a 'stopped' event (with reason 'breakpoint') is generated.
type SetBreakpointsRequest struct {
	// Object containing arguments for the command.
	Arguments *SetBreakpointsArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetBreakpointsResponse Response to 'setBreakpoints' request.
// Returned is information about each breakpoint created by this request.
// This includes the actual code location and whether the breakpoint could be verified.
// The breakpoints returned are in the same order as the elements of the 'breakpoints'
// (or the deprecated 'lines') array in the arguments.
type SetBreakpointsResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetDataBreakpointsArguments Arguments for 'setDataBreakpoints' request.
type SetDataBreakpointsArguments struct {
	// The contents of this array replaces all existing data breakpoints. An empty array clears all data breakpoints.
	Breakpoints []*DataBreakpoint `json:"breakpoints,omitempty"`
}

// SetDataBreakpointsRequest SetDataBreakpoints request; value of command field is 'setDataBreakpoints'.
// Replaces all existing data breakpoints with new data breakpoints.
// To clear all data breakpoints, specify an empty array.
// When a data breakpoint is hit, a 'stopped' event (with reason 'data breakpoint') is generated.
type SetDataBreakpointsRequest struct {
	// Object containing arguments for the command.
	Arguments *SetDataBreakpointsArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetDataBreakpointsResponse Response to 'setDataBreakpoints' request.
// Returned is information about each breakpoint created by this request.
type SetDataBreakpointsResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetExceptionBreakpointsArguments Arguments for 'setExceptionBreakpoints' request.
type SetExceptionBreakpointsArguments struct {
	// Configuration options for selected exceptions.
	ExceptionOptions []*ExceptionOptions `json:"exceptionOptions,omitempty"`

	// IDs of checked exception options. The set of IDs is returned via the 'exceptionBreakpointFilters' capability.
	Filters []string `json:"filters,omitempty"`
}

// SetExceptionBreakpointsRequest SetExceptionBreakpoints request; value of command field is 'setExceptionBreakpoints'.
// The request configures the debuggers response to thrown exceptions. If an exception is configured to break, a 'stopped' event is fired (with reason 'exception').
type SetExceptionBreakpointsRequest struct {
	// Object containing arguments for the command.
	Arguments *SetExceptionBreakpointsArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetExceptionBreakpointsResponse Response to 'setExceptionBreakpoints' request. This is just an acknowledgement, so no body field is required.
type SetExceptionBreakpointsResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetExpressionArguments Arguments for 'setExpression' request.
type SetExpressionArguments struct {
	// The l-value expression to assign to.
	Expression string `json:"expression,omitempty"`

	// Specifies how the resulting value should be formatted.
	Format *ValueFormat `json:"format,omitempty"`

	// Evaluate the expressions in the scope of this stack frame. If not specified, the expressions are evaluated in the global scope.
	FrameId float64 `json:"frameId,omitempty"`

	// The value expression to assign to the l-value expression.
	Value string `json:"value,omitempty"`
}

// SetExpressionRequest SetExpression request; value of command field is 'setExpression'.
// Evaluates the given 'value' expression and assigns it to the 'expression' which must be a modifiable l-value.
// The expressions have access to any variables and arguments that are in scope of the specified frame.
type SetExpressionRequest struct {
	// Object containing arguments for the command.
	Arguments *SetExpressionArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetExpressionResponse Response to 'setExpression' request.
type SetExpressionResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetFunctionBreakpointsArguments Arguments for 'setFunctionBreakpoints' request.
type SetFunctionBreakpointsArguments struct {
	// The function names of the breakpoints.
	Breakpoints []*FunctionBreakpoint `json:"breakpoints,omitempty"`
}

// SetFunctionBreakpointsRequest SetFunctionBreakpoints request; value of command field is 'setFunctionBreakpoints'.
// Replaces all existing function breakpoints with new function breakpoints.
// To clear all function breakpoints, specify an empty array.
// When a function breakpoint is hit, a 'stopped' event (with reason 'function breakpoint') is generated.
type SetFunctionBreakpointsRequest struct {
	// Object containing arguments for the command.
	Arguments *SetFunctionBreakpointsArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetFunctionBreakpointsResponse Response to 'setFunctionBreakpoints' request.
// Returned is information about each breakpoint created by this request.
type SetFunctionBreakpointsResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetVariableArguments Arguments for 'setVariable' request.
type SetVariableArguments struct {
	// Specifies details on how to format the response value.
	Format *ValueFormat `json:"format,omitempty"`

	// The name of the variable in the container.
	Name string `json:"name,omitempty"`

	// The value of the variable.
	Value string `json:"value,omitempty"`

	// The reference of the variable container.
	VariablesReference float64 `json:"variablesReference,omitempty"`
}

// SetVariableRequest SetVariable request; value of command field is 'setVariable'.
// Set the variable with the given name in the variable container to a new value.
type SetVariableRequest struct {
	// Object containing arguments for the command.
	Arguments *SetVariableArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SetVariableResponse Response to 'setVariable' request.
type SetVariableResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Source A Source is a descriptor for source code. It is returned from the debug adapter as part of a StackFrame and it is used by clients when specifying breakpoints.
type Source struct {
	// Optional data that a debug adapter might want to loop through the client. The client should leave the data intact and persist it across sessions. The client should not interpret the data.
	AdapterData interface{} `json:"adapterData,omitempty"`

	// The checksums associated with this file.
	Checksums []*Checksum `json:"checksums,omitempty"`

	// The short name of the source. Every source returned from the debug adapter has a name. When sending a source to the debug adapter this name is optional.
	Name string `json:"name,omitempty"`

	// The (optional) origin of this source: possible values 'internal module', 'inlined content from source map', etc.
	Origin string `json:"origin,omitempty"`

	// The path of the source to be shown in the UI. It is only used to locate and load the content of the source if no sourceReference is specified (or its value is 0).
	Path string `json:"path,omitempty"`

	// An optional hint for how to present the source in the UI. A value of 'deemphasize' can be used to indicate that the source is not available or that it is skipped on stepping.
	PresentationHint string `json:"presentationHint,omitempty"`

	// If sourceReference > 0 the contents of the source must be retrieved through the SourceRequest (even if a path is specified). A sourceReference is only valid for a session, so it must not be used to persist a source. The value should be less than or equal to 2147483647 (2^31 - 1).
	SourceReference float64 `json:"sourceReference,omitempty"`

	// An optional list of sources that are related to this source. These may be the source that generated this source.
	Sources []*Source `json:"sources,omitempty"`
}

// SourceArguments Arguments for 'source' request.
type SourceArguments struct {
	// Specifies the source content to load. Either source.path or source.sourceReference must be specified.
	Source *Source `json:"source,omitempty"`

	// The reference to the source. This is the same as source.sourceReference. This is provided for backward compatibility since old backends do not understand the 'source' attribute.
	SourceReference float64 `json:"sourceReference,omitempty"`
}

// SourceBreakpoint Properties of a breakpoint or logpoint passed to the setBreakpoints request.
type SourceBreakpoint struct {
	// An optional source column of the breakpoint.
	Column float64 `json:"column,omitempty"`

	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`

	// The source line of the breakpoint or logpoint.
	Line float64 `json:"line,omitempty"`

	// If this attribute exists and is non-empty, the backend must not 'break' (stop) but log the message instead. Expressions within {} are interpolated.
	LogMessage string `json:"logMessage,omitempty"`
}

// SourceRequest Source request; value of command field is 'source'.
// The request retrieves the source code for a given source reference.
type SourceRequest struct {
	// Object containing arguments for the command.
	Arguments *SourceArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// SourceResponse Response to 'source' request.
type SourceResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StackFrame A Stackframe contains the source location.
type StackFrame struct {
	// The column within the line. If source is null or doesn't exist, column is 0 and must be ignored.
	Column float64 `json:"column,omitempty"`

	// An optional end column of the range covered by the stack frame.
	EndColumn float64 `json:"endColumn,omitempty"`

	// An optional end line of the range covered by the stack frame.
	EndLine float64 `json:"endLine,omitempty"`

	// An identifier for the stack frame. It must be unique across all threads. This id can be used to retrieve the scopes of the frame with the 'scopesRequest' or to restart the execution of a stackframe.
	Id float64 `json:"id,omitempty"`

	// Optional memory reference for the current instruction pointer in this frame.
	InstructionPointerReference string `json:"instructionPointerReference,omitempty"`

	// The line within the file of the frame. If source is null or doesn't exist, line is 0 and must be ignored.
	Line float64 `json:"line,omitempty"`

	// The module associated with this frame, if any.
	ModuleId interface{} `json:"moduleId,omitempty"`

	// The name of the stack frame, typically a method name.
	Name string `json:"name,omitempty"`

	// An optional hint for how to present this frame in the UI. A value of 'label' can be used to indicate that the frame is an artificial frame that is used as a visual label or separator. A value of 'subtle' can be used to change the appearance of a frame in a 'subtle' way.
	PresentationHint string `json:"presentationHint,omitempty"`

	// The optional source of the frame.
	Source *Source `json:"source,omitempty"`
}

// StackFrameFormat Provides formatting information for a stack frame.
type StackFrameFormat struct {
	// Display the value in hex.
	Hex bool `json:"hex,omitempty"`

	// Includes all stack frames, including those the debug adapter might otherwise hide.
	IncludeAll bool `json:"includeAll,omitempty"`

	// Displays the line number of the stack frame.
	Line bool `json:"line,omitempty"`

	// Displays the module of the stack frame.
	Module bool `json:"module,omitempty"`

	// Displays the names of parameters for the stack frame.
	ParameterNames bool `json:"parameterNames,omitempty"`

	// Displays the types of parameters for the stack frame.
	ParameterTypes bool `json:"parameterTypes,omitempty"`

	// Displays the values of parameters for the stack frame.
	ParameterValues bool `json:"parameterValues,omitempty"`

	// Displays parameters for the stack frame.
	Parameters bool `json:"parameters,omitempty"`
}

// StackTraceArguments Arguments for 'stackTrace' request.
type StackTraceArguments struct {
	// Specifies details on how to format the stack frames.
	Format *StackFrameFormat `json:"format,omitempty"`

	// The maximum number of frames to return. If levels is not specified or 0, all frames are returned.
	Levels float64 `json:"levels,omitempty"`

	// The index of the first frame to return; if omitted frames start at 0.
	StartFrame float64 `json:"startFrame,omitempty"`

	// Retrieve the stacktrace for this thread.
	ThreadId float64 `json:"threadId,omitempty"`
}

// StackTraceRequest StackTrace request; value of command field is 'stackTrace'.
// The request returns a stacktrace from the current execution state.
type StackTraceRequest struct {
	// Object containing arguments for the command.
	Arguments *StackTraceArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StackTraceResponse Response to 'stackTrace' request.
type StackTraceResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StepBackArguments Arguments for 'stepBack' request.
type StepBackArguments struct {
	// Execute 'stepBack' for this thread.
	ThreadId float64 `json:"threadId,omitempty"`
}

// StepBackRequest StepBack request; value of command field is 'stepBack'.
// The request starts the debuggee to run one step backwards.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'step') after the step has completed. Clients should only call this request if the capability 'supportsStepBack' is true.
type StepBackRequest struct {
	// Object containing arguments for the command.
	Arguments *StepBackArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StepBackResponse Response to 'stepBack' request. This is just an acknowledgement, so no body field is required.
type StepBackResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StepInArguments Arguments for 'stepIn' request.
type StepInArguments struct {
	// Optional id of the target to step into.
	TargetId float64 `json:"targetId,omitempty"`

	// Execute 'stepIn' for this thread.
	ThreadId float64 `json:"threadId,omitempty"`
}

// StepInRequest StepIn request; value of command field is 'stepIn'.
// The request starts the debuggee to step into a function/method if possible.
// If it cannot step into a target, 'stepIn' behaves like 'next'.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'step') after the step has completed.
// If there are multiple function/method calls (or other targets) on the source line,
// the optional argument 'targetId' can be used to control into which target the 'stepIn' should occur.
// The list of possible targets for a given source line can be retrieved via the 'stepInTargets' request.
type StepInRequest struct {
	// Object containing arguments for the command.
	Arguments *StepInArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StepInResponse Response to 'stepIn' request. This is just an acknowledgement, so no body field is required.
type StepInResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StepInTarget A StepInTarget can be used in the 'stepIn' request and determines into which single target the stepIn request should step.
type StepInTarget struct {
	// Unique identifier for a stepIn target.
	Id float64 `json:"id,omitempty"`

	// The name of the stepIn target (shown in the UI).
	Label string `json:"label,omitempty"`
}

// StepInTargetsArguments Arguments for 'stepInTargets' request.
type StepInTargetsArguments struct {
	// The stack frame for which to retrieve the possible stepIn targets.
	FrameId float64 `json:"frameId,omitempty"`
}

// StepInTargetsRequest StepInTargets request; value of command field is 'stepInTargets'.
// This request retrieves the possible stepIn targets for the specified stack frame.
// These targets can be used in the 'stepIn' request.
// The StepInTargets may only be called if the 'supportsStepInTargetsRequest' capability exists and is true.
type StepInTargetsRequest struct {
	// Object containing arguments for the command.
	Arguments *StepInTargetsArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StepInTargetsResponse Response to 'stepInTargets' request.
type StepInTargetsResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StepOutArguments Arguments for 'stepOut' request.
type StepOutArguments struct {
	// Execute 'stepOut' for this thread.
	ThreadId float64 `json:"threadId,omitempty"`
}

// StepOutRequest StepOut request; value of command field is 'stepOut'.
// The request starts the debuggee to run again for one step.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'step') after the step has completed.
type StepOutRequest struct {
	// Object containing arguments for the command.
	Arguments *StepOutArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StepOutResponse Response to 'stepOut' request. This is just an acknowledgement, so no body field is required.
type StepOutResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// StoppedEvent Event message for 'stopped' event type.
// The event indicates that the execution of the debuggee has stopped due to some condition.
// This can be caused by a break point previously set, a stepping action has completed, by executing a debugger statement etc.
type StoppedEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// TerminateArguments Arguments for 'terminate' request.
type TerminateArguments struct {
	// A value of true indicates that this 'terminate' request is part of a restart sequence.
	Restart bool `json:"restart,omitempty"`
}

// TerminateRequest Terminate request; value of command field is 'terminate'.
// The 'terminate' request is sent from the client to the debug adapter in order to give the debuggee a chance for terminating itself.
type TerminateRequest struct {
	// Object containing arguments for the command.
	Arguments *TerminateArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// TerminateResponse Response to 'terminate' request. This is just an acknowledgement, so no body field is required.
type TerminateResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// TerminateThreadsArguments Arguments for 'terminateThreads' request.
type TerminateThreadsArguments struct {
	// Ids of threads to be terminated.
	ThreadIds []float64 `json:"threadIds,omitempty"`
}

// TerminateThreadsRequest TerminateThreads request; value of command field is 'terminateThreads'.
// The request terminates the threads with the given ids.
type TerminateThreadsRequest struct {
	// Object containing arguments for the command.
	Arguments *TerminateThreadsArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// TerminateThreadsResponse Response to 'terminateThreads' request. This is just an acknowledgement, so no body field is required.
type TerminateThreadsResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body interface{} `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// TerminatedEvent Event message for 'terminated' event type.
// The event indicates that debugging of the debuggee has terminated. This does **not** mean that the debuggee itself has exited.
type TerminatedEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Thread A Thread
type Thread struct {
	// Unique identifier for the thread.
	Id float64 `json:"id,omitempty"`

	// A name of the thread.
	Name string `json:"name,omitempty"`
}

// ThreadEvent Event message for 'thread' event type.
// The event indicates that a thread has started or exited.
type ThreadEvent struct {
	// Event-specific information.
	Body *Body `json:"body,omitempty"`

	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ThreadsRequest Threads request; value of command field is 'threads'.
// The request retrieves a list of all threads.
type ThreadsRequest struct {
	// Object containing arguments for the command.
	Arguments interface{} `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ThreadsResponse Response to 'threads' request.
type ThreadsResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// ValueFormat Provides formatting information for a value.
type ValueFormat struct {
	// Display the value in hex.
	Hex bool `json:"hex,omitempty"`
}

// Variable A Variable is a name/value pair.
// Optionally a variable can have a 'type' that is shown if space permits or when hovering over the variable's name.
// An optional 'kind' is used to render additional properties of the variable, e.g. different icons can be used to indicate that a variable is public or private.
// If the value is structured (has children), a handle is provided to retrieve the children with the VariablesRequest.
// If the number of named or indexed children is large, the numbers should be returned via the optional 'namedVariables' and 'indexedVariables' attributes.
// The client can use this optional information to present the children in a paged UI and fetch them in chunks.
type Variable struct {
	// Optional evaluatable name of this variable which can be passed to the 'EvaluateRequest' to fetch the variable's value.
	EvaluateName string `json:"evaluateName,omitempty"`

	// The number of indexed child variables.
	// The client can use this optional information to present the children in a paged UI and fetch them in chunks.
	IndexedVariables float64 `json:"indexedVariables,omitempty"`

	// Optional memory reference for the variable if the variable represents executable code, such as a function pointer.
	MemoryReference string `json:"memoryReference,omitempty"`

	// The variable's name.
	Name string `json:"name,omitempty"`

	// The number of named child variables.
	// The client can use this optional information to present the children in a paged UI and fetch them in chunks.
	NamedVariables float64 `json:"namedVariables,omitempty"`

	// Properties of a variable that can be used to determine how to render the variable in the UI.
	PresentationHint *VariablePresentationHint `json:"presentationHint,omitempty"`

	// The type of the variable's value. Typically shown in the UI when hovering over the value.
	Type string `json:"type,omitempty"`

	// The variable's value. This can be a multi-line text, e.g. for a function the body of a function.
	Value string `json:"value,omitempty"`

	// If variablesReference is > 0, the variable is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
	VariablesReference float64 `json:"variablesReference,omitempty"`
}

// VariablePresentationHint Optional properties of a variable that can be used to determine how to render the variable in the UI.
type VariablePresentationHint struct {
	// Set of attributes represented as an array of strings. Before introducing additional values, try to use the listed values.
	// Values:
	// 'static': Indicates that the object is static.
	// 'constant': Indicates that the object is a constant.
	// 'readOnly': Indicates that the object is read only.
	// 'rawString': Indicates that the object is a raw string.
	// 'hasObjectId': Indicates that the object can have an Object ID created for it.
	// 'canHaveObjectId': Indicates that the object has an Object ID associated with it.
	// 'hasSideEffects': Indicates that the evaluation had side effects.
	// etc.
	Attributes []string `json:"attributes,omitempty"`

	// The kind of variable. Before introducing additional values, try to use the listed values.
	// Values:
	// 'property': Indicates that the object is a property.
	// 'method': Indicates that the object is a method.
	// 'class': Indicates that the object is a class.
	// 'data': Indicates that the object is data.
	// 'event': Indicates that the object is an event.
	// 'baseClass': Indicates that the object is a base class.
	// 'innerClass': Indicates that the object is an inner class.
	// 'interface': Indicates that the object is an interface.
	// 'mostDerivedClass': Indicates that the object is the most derived class.
	// 'virtual': Indicates that the object is virtual, that means it is a synthetic object introduced by the adapter for rendering purposes, e.g. an index range for large arrays.
	// 'dataBreakpoint': Indicates that a data breakpoint is registered for the object.
	// etc.
	Kind string `json:"kind,omitempty"`

	// Visibility of variable. Before introducing additional values, try to use the listed values.
	// Values: 'public', 'private', 'protected', 'internal', 'final', etc.
	Visibility string `json:"visibility,omitempty"`
}

// VariablesArguments Arguments for 'variables' request.
type VariablesArguments struct {
	// The number of variables to return. If count is missing or 0, all variables are returned.
	Count float64 `json:"count,omitempty"`

	// Optional filter to limit the child variables to either named or indexed. If omitted, both types are fetched.
	Filter string `json:"filter,omitempty"`

	// Specifies details on how to format the Variable values.
	Format *ValueFormat `json:"format,omitempty"`

	// The index of the first variable to return; if omitted children start at 0.
	Start float64 `json:"start,omitempty"`

	// The Variable reference.
	VariablesReference float64 `json:"variablesReference,omitempty"`
}

// VariablesRequest Variables request; value of command field is 'variables'.
// Retrieves all child variables for the given variable reference.
// An optional filter can be used to limit the fetched children to either named or indexed children.
type VariablesRequest struct {
	// Object containing arguments for the command.
	Arguments *VariablesArguments `json:"arguments,omitempty"`

	// The command to execute.
	Command string `json:"command,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// VariablesResponse Response to 'variables' request.
type VariablesResponse struct {
	// Contains request result if success is true and optional error details if success is false.
	Body *Body `json:"body,omitempty"`

	// The command requested.
	Command string `json:"command,omitempty"`

	// Contains the raw error in short form if 'success' is false.
	// This raw error might be interpreted by the frontend and is not shown in the UI.
	// Some predefined values exist.
	// Values:
	// 'cancelled': request was cancelled.
	// etc.
	Message string `json:"message,omitempty"`

	// Sequence number of the corresponding request.
	RequestSeq float64 `json:"request_seq,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Outcome of the request.
	// If true, the request was successful and the 'body' attribute may contain the result of the request.
	// If the value is false, the attribute 'message' contains the error in short form and the 'body' may contain additional information (see 'ErrorResponse.body.error').
	Success bool `json:"success,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}
