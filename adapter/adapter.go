// Copyright 2020 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapter

// BeforeExitCallback Dispose and allow exit to continue normally
type BeforeExitCallback struct{}

// Body Event-specific information.
type Body struct {
	Reason   string  `json:"reason,omitempty"`
	ThreadId float64 `json:"threadId,omitempty"`
}

// Breakpoint
type Breakpoint struct {
	// If true breakpoint could be set (but not necessarily at the desired location).
	Verified bool `json:"verified,omitempty"`
}

// BreakpointEvent
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

// CapabilitiesEvent
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

// CompletionItem
type CompletionItem struct {
	// The label of this completion item. By default this is also the text that is inserted when selecting this completion.
	Label string `json:"label,omitempty"`

	// This value determines how many characters are overwritten by the completion text.
	// If missing the value 0 is assumed which results in the completion text being inserted.
	Length float64 `json:"length,omitempty"`

	// This value determines the location (in the CompletionsRequest's 'text' attribute) where the completion text is added.
	// If missing the text is added at the location specified by the CompletionsRequest's 'column' attribute.
	Start float64 `json:"start,omitempty"`
}

// ContinuedEvent
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

// DebugProtocolCapabilities Information about the capabilities of a debug adapter.
type DebugProtocolCapabilities struct {
	// The set of additional module information exposed by the debug adapter.
	AdditionalModuleColumns []*DebugProtocolColumnDescriptor `json:"additionalModuleColumns,omitempty"`

	// The set of characters that should trigger completion in a REPL. If not specified, the UI should assume the '.' character.
	CompletionTriggerCharacters []string `json:"completionTriggerCharacters,omitempty"`

	// Available filters or options for the setExceptionBreakpoints request.
	ExceptionBreakpointFilters []*DebugProtocolExceptionBreakpointsFilter `json:"exceptionBreakpointFilters,omitempty"`

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

// DebugProtocolColumnDescriptor A ColumnDescriptor specifies what module attribute to show in a column of the ModulesView, how to format it, and what the column's label should be.
// It is only used if the underlying UI actually supports this level of customization.
type DebugProtocolColumnDescriptor struct {
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

// DebugProtocolExceptionBreakpointsFilter An ExceptionBreakpointsFilter is shown in the UI as an option for configuring how exceptions are dealt with.
type DebugProtocolExceptionBreakpointsFilter struct {
	// Initial value of the filter. If not specified a value 'false' is assumed.
	Default bool `json:"default,omitempty"`

	// The internal ID of the filter. This value is passed to the setExceptionBreakpoints request.
	Filter string `json:"filter,omitempty"`

	// The name of the filter. This will be shown in the UI.
	Label string `json:"label,omitempty"`
}

// DebugProtocolMessage
type DebugProtocolMessage struct{}

// DebugSession
type DebugSession struct {
	ClientColumnsStartAt1   bool                        `json:"_clientColumnsStartAt1,omitempty"`
	ClientLinesStartAt1     bool                        `json:"_clientLinesStartAt1,omitempty"`
	ClientPathsAreURIs      bool                        `json:"_clientPathsAreURIs,omitempty"`
	ContentLength           float64                     `json:"_contentLength,omitempty"`
	DebuggerColumnsStartAt1 bool                        `json:"_debuggerColumnsStartAt1,omitempty"`
	DebuggerLinesStartAt1   bool                        `json:"_debuggerLinesStartAt1,omitempty"`
	DebuggerPathsAreURIs    bool                        `json:"_debuggerPathsAreURIs,omitempty"`
	Error                   *EmitterError               `json:"error,omitempty"`
	IsServer                bool                        `json:"_isServer,omitempty"`
	OnError                 *Event0Error                `json:"onError,omitempty"`
	OnSendMessage           *Event0DebugProtocolMessage `json:"onSendMessage,omitempty"`
	PendingRequests         interface{}                 `json:"_pendingRequests,omitempty"`

	// Raw data is stored in instances of the Buffer class.
	// A Buffer is similar to an array of integers but corresponds to a raw memory allocation outside the V8 heap.  A Buffer cannot be resized.
	// Valid string encodings: 'ascii'|'utf8'|'utf16le'|'ucs2'(alias of 'utf16le')|'base64'|'binary'(deprecated)|'hex'
	RawData        []float64                    `json:"_rawData,omitempty"`
	SendMessage    *EmitterDebugProtocolMessage `json:"sendMessage,omitempty"`
	Sequence       float64                      `json:"_sequence,omitempty"`
	WritableStream *NodeJSWritableStream        `json:"_writableStream,omitempty"`
}

// Disposable0
type Disposable0 struct{}

// Emitter
type Emitter struct {
	Event    *Event0T    `json:"_event,omitempty"`
	Listener *Listener   `json:"_listener,omitempty"`
	This     interface{} `json:"_this,omitempty"`
}

// EmitterDebugProtocolMessage
type EmitterDebugProtocolMessage struct {
	Event    *Event0DebugProtocolMessage `json:"_event,omitempty"`
	Listener *Listener                   `json:"_listener,omitempty"`
	This     interface{}                 `json:"_this,omitempty"`
}

// EmitterError
type EmitterError struct {
	Event    *Event0Error `json:"event,omitempty"`
	Listener *Listener    `json:"_listener,omitempty"`
	This     interface{}  `json:"_this,omitempty"`
}

// Event
type Event struct {
	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Event0
type Event0 struct{}

// Event0DebugProtocolMessage
type Event0DebugProtocolMessage struct{}

// Event0Error
type Event0Error struct{}

// Event0T
type Event0T struct{}

// Handles
type Handles struct {
	HandleMap   interface{} `json:"_handleMap,omitempty"`
	NextHandle  float64     `json:"_nextHandle,omitempty"`
	STARTHANDLE float64     `json:"START_HANDLE,omitempty"`
}

// IDisposable
type IDisposable struct{}

// IInternalLoggerOptions
type IInternalLoggerOptions struct {
	ConsoleMinLogLevel float64 `json:"consoleMinLogLevel,omitempty"`
	LogFilePath        string  `json:"logFilePath,omitempty"`
	PrependTimestamp   bool    `json:"prependTimestamp,omitempty"`
}

// ILogCallback
type ILogCallback struct{}

// ILogItem
type ILogItem struct {
	Level float64 `json:"level,omitempty"`
	Msg   string  `json:"msg,omitempty"`
}

// ILogger
type ILogger struct{}

// InitializedEvent
type InitializedEvent struct {
	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// InternalLogger Manages logging, whether to console.log, file, or VS Code console.
// Encapsulates the state specific to each logging session
type InternalLogger struct {
	// Dispose and allow exit to continue normally
	BeforeExitCallback *BeforeExitCallback `json:"beforeExitCallback,omitempty"`

	// Dispose and exit
	DisposeCallback interface{} `json:"disposeCallback,omitempty"`

	// Log info that meets minLogLevel is sent to this callback.
	LogCallback *LogCallback `json:"_logCallback,omitempty"`

	// Write steam for log file
	LogFileStream *WriteStream `json:"_logFileStream,omitempty"`
	LogToConsole  bool         `json:"_logToConsole,omitempty"`
	MinLogLevel   float64      `json:"_minLogLevel,omitempty"`

	// Whether to add a timestamp to messages in the logfile
	PrependTimestamp bool `json:"_prependTimestamp,omitempty"`
}

// Listener
type Listener struct{}

// LoadedSourceEvent
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

// LogCallback Log info that meets minLogLevel is sent to this callback.
type LogCallback struct{}

// LogOutputEvent
type LogOutputEvent struct {
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

// Logger
type Logger struct {
	CurrentLogger       *InternalLogger `json:"_currentLogger,omitempty"`
	LogFilePathFromInit string          `json:"_logFilePathFromInit,omitempty"`
	PendingLogQ         []*ILogItem     `json:"_pendingLogQ,omitempty"`
}

// LoggingDebugSession
type LoggingDebugSession struct {
	ClientColumnsStartAt1   bool                        `json:"_clientColumnsStartAt1,omitempty"`
	ClientLinesStartAt1     bool                        `json:"_clientLinesStartAt1,omitempty"`
	ClientPathsAreURIs      bool                        `json:"_clientPathsAreURIs,omitempty"`
	ContentLength           float64                     `json:"_contentLength,omitempty"`
	DebuggerColumnsStartAt1 bool                        `json:"_debuggerColumnsStartAt1,omitempty"`
	DebuggerLinesStartAt1   bool                        `json:"_debuggerLinesStartAt1,omitempty"`
	DebuggerPathsAreURIs    bool                        `json:"_debuggerPathsAreURIs,omitempty"`
	Error                   *EmitterError               `json:"error,omitempty"`
	IsServer                bool                        `json:"_isServer,omitempty"`
	ObsoleteLogFilePath     string                      `json:"obsolete_logFilePath,omitempty"`
	OnError                 *Event0Error                `json:"onError,omitempty"`
	OnSendMessage           *Event0DebugProtocolMessage `json:"onSendMessage,omitempty"`
	PendingRequests         interface{}                 `json:"_pendingRequests,omitempty"`

	// Raw data is stored in instances of the Buffer class.
	// A Buffer is similar to an array of integers but corresponds to a raw memory allocation outside the V8 heap.  A Buffer cannot be resized.
	// Valid string encodings: 'ascii'|'utf8'|'utf16le'|'ucs2'(alias of 'utf16le')|'base64'|'binary'(deprecated)|'hex'
	RawData        []float64                    `json:"_rawData,omitempty"`
	SendMessage    *EmitterDebugProtocolMessage `json:"sendMessage,omitempty"`
	Sequence       float64                      `json:"_sequence,omitempty"`
	WritableStream *NodeJSWritableStream        `json:"_writableStream,omitempty"`
}

// Message
type Message struct {
	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Module
type Module struct {
	// Unique identifier for the module.
	Id interface{} `json:"id,omitempty"`

	// A name of the module.
	Name string `json:"name,omitempty"`
}

// ModuleEvent
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

// NodeJSWritableStream
type NodeJSWritableStream struct {
	Writable bool `json:"writable,omitempty"`
}

// OutputEvent
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

// ProtocolServer
type ProtocolServer struct {
	ContentLength   float64                     `json:"_contentLength,omitempty"`
	Error           *EmitterError               `json:"error,omitempty"`
	OnError         *Event0Error                `json:"onError,omitempty"`
	OnSendMessage   *Event0DebugProtocolMessage `json:"onSendMessage,omitempty"`
	PendingRequests interface{}                 `json:"_pendingRequests,omitempty"`

	// Raw data is stored in instances of the Buffer class.
	// A Buffer is similar to an array of integers but corresponds to a raw memory allocation outside the V8 heap.  A Buffer cannot be resized.
	// Valid string encodings: 'ascii'|'utf8'|'utf16le'|'ucs2'(alias of 'utf16le')|'base64'|'binary'(deprecated)|'hex'
	RawData        []float64                    `json:"_rawData,omitempty"`
	SendMessage    *EmitterDebugProtocolMessage `json:"sendMessage,omitempty"`
	Sequence       float64                      `json:"_sequence,omitempty"`
	WritableStream *NodeJSWritableStream        `json:"_writableStream,omitempty"`
}

// Response
type Response struct {
	// The command requested.
	Command string `json:"command,omitempty"`

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

// Scope
type Scope struct {
	// If true, the number of variables in this scope is large or expensive to retrieve.
	Expensive bool `json:"expensive,omitempty"`

	// Name of the scope such as 'Arguments', 'Locals', or 'Registers'. This string is shown in the UI as is and can be translated.
	Name string `json:"name,omitempty"`

	// The variables of this scope can be retrieved by passing the value of variablesReference to the VariablesRequest.
	VariablesReference float64 `json:"variablesReference,omitempty"`
}

// Source
type Source struct {
	// The short name of the source. Every source returned from the debug adapter has a name. When sending a source to the debug adapter this name is optional.
	Name string `json:"name,omitempty"`

	// The path of the source to be shown in the UI. It is only used to locate and load the content of the source if no sourceReference is specified (or its value is 0).
	Path string `json:"path,omitempty"`

	// If sourceReference > 0 the contents of the source must be retrieved through the SourceRequest (even if a path is specified). A sourceReference is only valid for a session, so it must not be used to persist a source. The value should be less than or equal to 2147483647 (2^31 - 1).
	SourceReference float64 `json:"sourceReference,omitempty"`
}

// StackFrame
type StackFrame struct {
	// The column within the line. If source is null or doesn't exist, column is 0 and must be ignored.
	Column float64 `json:"column,omitempty"`

	// An identifier for the stack frame. It must be unique across all threads. This id can be used to retrieve the scopes of the frame with the 'scopesRequest' or to restart the execution of a stackframe.
	Id float64 `json:"id,omitempty"`

	// The line within the file of the frame. If source is null or doesn't exist, line is 0 and must be ignored.
	Line float64 `json:"line,omitempty"`

	// The name of the stack frame, typically a method name.
	Name string `json:"name,omitempty"`

	// The optional source of the frame.
	Source *Source `json:"source,omitempty"`
}

// StoppedEvent
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

// TerminatedEvent
type TerminatedEvent struct {
	// Type of event.
	Event string `json:"event,omitempty"`

	// Sequence number (also known as message ID). For protocol messages of type 'request' this ID can be used to cancel the request.
	Seq float64 `json:"seq,omitempty"`

	// Message type.
	// Values: 'request', 'response', 'event', etc.
	Type string `json:"type,omitempty"`
}

// Thread
type Thread struct {
	// Unique identifier for the thread.
	Id float64 `json:"id,omitempty"`

	// A name of the thread.
	Name string `json:"name,omitempty"`
}

// ThreadEvent
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

// VSCodeDebugAdapter A structurally equivalent copy of vscode.DebugAdapter
type VSCodeDebugAdapter struct {
	OnError       *Event0Error                `json:"onError,omitempty"`
	OnSendMessage *Event0DebugProtocolMessage `json:"onSendMessage,omitempty"`
}

// Variable
type Variable struct {
	// The variable's name.
	Name string `json:"name,omitempty"`

	// The variable's value. This can be a multi-line text, e.g. for a function the body of a function.
	Value string `json:"value,omitempty"`

	// If variablesReference is > 0, the variable is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
	VariablesReference float64 `json:"variablesReference,omitempty"`
}

// WriteStream
type WriteStream struct {
	BytesWritten float64     `json:"bytesWritten,omitempty"`
	Path         interface{} `json:"path,omitempty"`
	Writable     bool        `json:"writable,omitempty"`
}
