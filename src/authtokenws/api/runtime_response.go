package api

//
// RuntimeResponse -- response structure to the runtime request
//
type RuntimeResponse struct {
   Version         string `json:"version"`
   CPUCount        int    `json:"cpus"`
   GoRoutineCount  int    `json:"go_routines"`
   ObjectCount     uint64 `json:"heap_objects"`
   AllocatedMemory uint64 `json:"allocated_mem"`
}

//
// end of file
//
