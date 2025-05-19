// Here’s a simplified example of prioritizing one channel over another using nested *select*:
for {
    select {
    case val := <-highPriorityChan:
        // Always handle high-priority messages first
        processHigh(val)
    default:
        select {
        case val := <-lowPriorityChan:
            processLow(val)
        case <-time.After(time.Second):
            // Optional: timeout or periodic check
        }
    }
}
