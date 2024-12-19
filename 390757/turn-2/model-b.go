package main
import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type SensorData struct {
    SensorID string
    Timestamp time.Time
    Value    float64
}

type SensorSimulator struct {
    SensorID     string
    SampleRate   time.Duration // Duration between samples
    DataCh      chan SensorData
    quitCh      chan struct{}
    stopOnce    sync.Once
    wg          sync.WaitGroup
}

func (s *SensorSimulator) Start() {
    s.wg.Add(1)
    defer s.wg.Done()
    ticker := time.NewTicker(s.SampleRate)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            s.simulateSample()
        case <-s.quitCh:
            return
        }
    }
}

func (s *SensorSimulator) simulateSample() {
    data := SensorData{
        SensorID: s.SensorID,
        Timestamp: time.Now(),
        Value:     rand.Float64(), // Generate random value for this example
    }
    s.DataCh <- data
}

// Stop the sensor simulation.
func (s *SensorSimulator) Stop() {
    s.stopOnce.Do(func() {
        close(s.quitCh)
    })
}

type EventProcessor struct {
    DataCh    chan SensorData
    MaxQueueLen int
    wg        sync.WaitGroup
}

func NewEventProcessor(maxQueueLen int) *EventProcessor {
    return &EventProcessor{
        DataCh:    make(chan SensorData, maxQueueLen),
        MaxQueueLen: maxQueueLen,
    }
}

func (ep *EventProcessor) Run() {
    ep.wg.Add(1)
    defer ep.wg.Done()
    for data := range ep.DataCh {
        // Process the event here. For demonstration, we'll just print the data.
        fmt.Printf("Event processed: %+v\n", data)
    }
}

// AddData adds a new data event to the event processor.
func (ep *EventProcessor) AddData(data SensorData) bool {
    select {
    case ep.DataCh <- data:
        return true // Data added successfully
    default:
        // Channel is full, event dropped. We can implement different strategies for handling dropped events.
        fmt.Println("Event dropped due to full queue")
        return false
    }
}

func (ep *EventProcessor) Stop() {
    close(ep.DataCh)
    ep.wg.Wait()
}

func main() {
    numSensors := 10 // Number of sensor simulators
    sensorSampleRate := 1 * time.Second // Sample rate for each sensor
    maxQueueLen := 1000  // Maximum length of the event queue

    ep := NewEventProcessor(maxQueueLen)
    go ep.Run()

    // Create and start sensor simulators
    var simulators []*SensorSimulator
    for i := 0; i < numSensors; i++ {
        sim := &SensorSimulator{
            SensorID:   fmt.Sprintf("Sensor-%d", i+1),
            SampleRate: sensorSampleRate,
            DataCh:     make(chan SensorData, 10), // Buffered channel for each sensor
            quitCh:     make(chan struct{}),
        }
        simulators = append(simulators, sim)
        go sim.Start()
    }

    // Main loop to add data events from sensors to the event processor.
    for {
        time.Sleep(time.Second)
        for _, sim := range simulators {
            data, more := <-sim.DataCh
            if more {
