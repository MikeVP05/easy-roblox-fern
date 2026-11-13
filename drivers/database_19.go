package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.2" )

func u9GUT0uhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SDJJHiGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGwkaKzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jK356GTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDUdZbsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ex3upEJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiIQULHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6WXEddhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNBTtXnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgXfWCbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7IeqUbRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gz5l1ybmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nmw3Gv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xMzHcGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ve78fg7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6MAUSBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTpPwyVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Myr0hkpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxUP7M6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmV03MKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func naDC8aCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Kjdj5EVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7OlEZj3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wh7ij8N0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WYGh7MkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1kpfr4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PjmDxuPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNd3mmchWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJSSNl9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vMG15LNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bv8aRe8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dq3ipTfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59VIoYEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbUCbOsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cMSAKkOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2r3vXf1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pFlAfrRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwMg9MsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eB50pWyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wiTR9G0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X80GXcNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCZy66KpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjZoaNzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIQZDeV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2mtaeCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LchWinmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTzpMznqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDc1YB9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spaHjDLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zrd0ArMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbIo9bSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqaP2sIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7sfUQ5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDjkch1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsNTtIsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjyL7XamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q3fJEDXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IIbt2p2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8iYkSrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4sQuYNAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6QTK2vLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnawDehdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtQbu7aqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gPYfqZTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BEMSTqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtR9b9THWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHpveZqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3kEIsUwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func no6H7E7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ni7TywnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5N3LUB65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4XerJUGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHuUVdxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jriFCrM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JWJmutyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddMzyziRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrlsqQ8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uT3aMfHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9r76A8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jO1vLD17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wOhTgi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mkciuMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zheVbGGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lMgzL4INWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIKdjpkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUbKM5X7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cRxSPzHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggVHgUh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fm55yLw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EduhTqzpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZxPlH1CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uk8nTGCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRaYlGOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuJISKnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7j12I2eKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxUjl8bBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpGY9TaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZPr5cDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1GFU1F3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hE1UdsKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6PdxFa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1YPnvPwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2P6MLWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fvkItjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NORe3b0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYpTUtqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTbkCdrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrhJ334eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfcukU6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKVL5j5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dUtkdBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmPXX87wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1P8NQc4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfBUThWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8D91JCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HiUYXQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xW2K9nTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQsrtoSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5sBEwYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2GnufQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vp7uLqEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9oXwf0rnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R9x04ZedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUUfmVZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDcTDzrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3FyjmXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g9RbycpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRah4IedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adVn49OyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 809geYRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfuwXPsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYImUVE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQpQxyuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWPYrGADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYLjYS8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvEbGM9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPfgJPmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZSkplvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qYgRgmLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35BS9UwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCdZIVySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCjAo0OZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQMv4hPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g883zaKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AlT0pOa3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EoCr0d1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZidnUeULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogIHWqegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IWG2Cy1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4izK6TKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QlrzaINrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nIalbWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgJ6zgiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ui05SPblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1Od19W7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1xFJbkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCGwsJvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7O5VunAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K9tHJtPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1YQ8ToXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKYHtsLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZ2Pv8twWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SmIj6mycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwVoyKUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKVuTgagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2AklANUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FhVnrOXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpw4KcqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHbpm9ZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njMR401TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5aFboWrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToVUtMwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2MxiqsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2z3MwFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t37qKLCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8vjPpdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtFmQXy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 97izvjwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FhQiuT4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q46KDdX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlXbLijPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZjmfXWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXrHSq4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0J2WJqvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SiaPna3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func An6odJ6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYOLYjGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MjMcbjY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtmYRFqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hs2FDarnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V35SNx1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38XpDiIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IaS4bi3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IsTQldVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOyq8PlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Md1nlNgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3OVhFAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AkkfJ5jmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1tfKGcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vi0AysVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVc7hTtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlMFgU87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 424BOvaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXr5fL1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2ermJwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdNQ0PWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqcQZTUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nL5UFs7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdGs2PBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1EiJtNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xe3CgbMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miRf30veWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QcZ5KBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CetObC88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pisn22gDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9v7EqYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8avzXLXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6tpVIVGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzjvhCsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3WrYkUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwmFpfZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iX8W222IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soDHqlE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zu3xwOU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8vNwT9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuLE6fGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IaSWbNM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLAjS8feWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSIZHAFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Rp5YE2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TURkg6V0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MmPyVZoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0WnIpshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cC0x5ZYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVjMRvmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zn3OtpPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXl6IUSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ME9rwtROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQKGDaLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJ56wykFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hV4w2qWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyrTxcaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6DpKK9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXd4xaIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SsbT2HmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Et0VfIQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8rA4a2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPulqKIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G08hlKXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Gn7kn1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovhERjNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPL9uGblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jsvgh1lTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uICBYyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Cpbnee6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bld6q6JGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqhrnxbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bo876q2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ex0dhhhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFPYklFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQS1m8xTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSkjbbVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQo8xaHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wA2zpr6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5bvuO39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D5eC2ntvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTHEd7uiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CFbkfnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuG6GNsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yIqGEgj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOImmNSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWowRQ6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EA2rioqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDblEpbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wi7lrpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGvmjm16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gREHV3oMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHRPGmdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQZB4bGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9b2sJVbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mnfQ3DCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNQqPfPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iirgceRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76uHMRi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVpwwyX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tR7hEK3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmmKjGhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CUoirFKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8IgojdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6ZYEabHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ilgue15WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rk1z41cIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqwUFKGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwHQHwxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 011Xg9koWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgO2VtRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEMJdyrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hahxu373Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbCAYCiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWBmK6b1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJ5rpSrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqFZAVEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCrIcUv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EiYgwJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aeFhPlyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMN6Gv0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BuiAh7dmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrorz1yMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xoWwWhhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1M7ouqBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TdMrWuwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ZNGaR1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbMudNSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdEWk8uWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLyspovQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v37sxluvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImymXCK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCYvS5gwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzfuuhyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQh2eUUxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzRveI6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgdZPLJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZwCVMLfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHqRoUXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8p82aY6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRLzwyYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvsLAOIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqeSxZJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dpikd5blWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JcT7qkFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39O2RKGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umFWYMb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgKPLk0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6RsXsMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63fgUbwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsylDVn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGcAHPTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XyDW9WBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vp3HlYjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pICK1hZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZjDVItRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8GhJB7MWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkpPdRg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNS0Jyi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLAYsv1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LPsFHBCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FKGYBUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XEUMa9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOAI6AzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3gA7FIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rzmkCEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9DBaBjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nf7GRxrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tENJXBsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upqzPcEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLAVUNIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59oXRy4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9GiDp7asWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCeuTBBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8of7vZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDmfD7ISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpbKmxBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmF9KKwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EpQpJqfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXSzc1BaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agq84ZfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWZbrtXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZDgajzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICi6hJ2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgmnuuCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LOlvlaTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNQmf3EIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViG6sEciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3lxxoOAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LxWAG2KvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRa1RJqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Qsst83IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDmUFWGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8fb7N1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6d0aG5YOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKtaSbRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2Se91BKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5U5zToYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vdo0klmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nh8NpzKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUuLhTU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsI7tvcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LhxvqgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ISR0XgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWZqsWO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfcLp5CfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVcpYVtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UN5fa4gUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3hMny9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h09s5cJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l27keVyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUgIz9yjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zVCXB8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TlL4iycrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8IkLe3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mg83IHfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juLERB12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzUJKd0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1tPhSQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QljVLwM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiJW9AQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vt1ffziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JtdJPIzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mv0AtAhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func alK0psrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8zIQj8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLeRgO24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fb9pHrOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func So9d1hDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ue8xb2ZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qjl0VM0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VG2aJAxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axNcDHgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afWujrG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdIxBKIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7v74wsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rxkzr9v1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQj0WP0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ovo5yAxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpsqS4cwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJ1du32GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPYB8s8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9EIkPlLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkv0TYAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgRT0NOhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nE5pmcDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdZv88xSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nJidebPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4i5MIyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2nmNmtAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzOz5hYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PrOCU4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuLyNjvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dds9DARQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLXFLk5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5vj2UHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LVtqemJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bapxj5EWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xiVKmYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PghKneOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmqoBobzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 02Yd9I1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFDoUfMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjbocpRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func In606jB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DiEillLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWGe24JOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5vdEurhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4NKyo1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8sMKn3JwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5YoQTI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N7cwlNDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZkkgeGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XK6UbJihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6p7A5LxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpnHuFTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Fsu4PSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GlHOSE3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j71BuBr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRdhnODRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3WSWqh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdX0PIHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6AoVBr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNWXgEHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiuSrdPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iAI0u2LvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkvAuM1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZXUtw1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGXMWAsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RxHn9rYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oTmn9whtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLYWrD8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYZig7kjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qv9VQ4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEMeqByWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G83OfQ3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2VjAWKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yaJxcTbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVDYZNVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWHT0T4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tr7zvC69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inZfHoLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdLCBRI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3NBo4fLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qy29MddoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGFqZidcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhZ7DBvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjjOXlptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GkiJgUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func turO0Ds6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9pWZpmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1ezVi4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdMC7ciVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qu9rEFzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjYYLZUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vk28Eb1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AeWpo9GvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjQRwT1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FKWCusSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J4gdSYGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAvOdBEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUzueoc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7EipT7TKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func maHXMhwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysxMMKtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8upULO5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L6W0qxm2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MgeBOiLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFaGKNsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmf7vq8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROQ6yywWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oT8ZdMUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jk7WusQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qo7XRaahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOcJGgqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KuoingDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSjRP1HdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJzbBNWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsdmJDWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylfQBUe3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlIIaWb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27V5uMOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHA5q29RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOmh16UwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dp1OM5tzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dyrdKnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vixuKvTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2s5kW4SEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpvJAcWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGWdvqr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIXorbOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JsWuMsY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fdmrng54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gy2pMy8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6GrPKctFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGQR2ExyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diBHZZGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LebndoxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igLEj50tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L6KpsAxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdVrNpaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1WnT28PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uR4Nie4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyqdJ1g8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvUqckVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dRNUZhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLMe3druWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RreBnQOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIIQKOC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFufjb69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCyJeA7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoUhz6nNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhfVRhkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XNZGnh7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuPZKXKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LEOJHNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbIuNq6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3pEF4hoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofhI7fWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFL0WKbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Il5xU0ISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxJO5TOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WruZic8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkiWrELyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aj4fMOiTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nddJp4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FeoSAAOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8NvwA65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqLjDR2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSnFoPYdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRrqcpsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HjrhlM3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmyY1X5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsYela4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vvy5nnWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hb1MPFBBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GcD4fcguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRIPTOuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFxL9aU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giR0x24DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDXWaMAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0Qdq06VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPB74l5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0ep92LnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JrUUquNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzODNdy8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wh9RYvbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBEfrLYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53WrO082Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OeYUJlXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGLQl66zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXcF2CthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKD1JsArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uz3xDPCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuwbgPInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WRPmyTGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhdCtIQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0bmsMqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3rJty4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80UPzYVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnKWpQUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOcrlTZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hycwJAmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3gPEsPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oc65K8H4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZnOmitQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJnPAly4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twyOT8XYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zyOLHnitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0xzICIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func di0kaSiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DanYlOasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1BgOmyf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpE2jAUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYvSlvEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0UI6dy8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxI0Aay6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l97EtIMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clJVGAJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqMgD4MHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o7K7X40HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpoWKXThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jo2cI1r6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQhSmZhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ja4oLnS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPDnoNDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLYQ5uT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oumUmJc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9K8oNxeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DAAQFajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwftfCpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFJe4ZNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvTI60SAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4VcQQgRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuAzYzfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M38qf1e1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7qWJu98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nm3DQtwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUlvjhN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKMOzUcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eK8MSnFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uO1hySndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWmamAi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Th2ghbzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwMTtwFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 00FQ4uBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRgOcP6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OeIvAvQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AyeHllSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOlEVEC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyb2WVwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llsyGMANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvKcMAXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cYpI3ndsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lKKi79DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AN5oJc5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXB9sB8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWMlUBydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4XGpPDNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqphjd7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AopFrcMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ntREVTxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHfw7ISrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i17LOodIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNH6baRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Fc0T7LLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2uAuC4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lq9twe5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IhXld1W4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnYh1Pi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxv0jSacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZixQVPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q97UjtiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqX6cJjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6OFDzkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhZmlTQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xWCqBynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8oh990iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGimfqpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5y0U8mMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ze1nfbmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jc3iwzIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1U0VRlIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88XOLIaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rfxu500Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90jjXKdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f78W4rgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugFVcUJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cq29qRWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekvBq4P1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cn0YSTgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWS6FHPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PXzVEPH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MuFUMdygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhDYxLwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjF0O0fHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTy4G8KUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhJ67uorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5cBtDwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFG0YHhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2fhrLloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWAiOAX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJx1rlwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOwt5my1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6GixghdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNCkqod5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BfmJh1zaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBjzSVacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBrm7XKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsRY0GllWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hOLC7utRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcbkRbzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfVatlS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2itumLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLgm9TrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXq4K1aqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qie5doNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4mWqowhBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Wa3hTDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6eM1UhmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ns96BzosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UQurWHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFpNDzBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOBIPTALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vrxv5eJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oa0Yz2wYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPRMv1IdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KY1iFIOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2mC70vlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q3LF1125Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K9uG5E0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ssm3IGraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UU79nojkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8toI9yQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uOLzvM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YwvbefgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ABx48FpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L12NbOfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHzIgFODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIEiKuEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcfYpGC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wbekWDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USurxAZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgjVh1TZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHJ5hjb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzVU1f6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ow31H8UhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5CSrZWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmSeMg81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjQXmUzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yKfS37gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvusrTQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uA4ZaOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nO3ZtX7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDF90VUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hKFoPHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func saqZssVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3uzSNPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwrfhzOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fAfmA7JTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCxtWScKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9GuKUH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usXfSqIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLqPAPhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func slULctJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAm0nzqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVUzoMdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xc9QG9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4RaT3UGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLOmk72LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQy4YPPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xeMviZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHjrAy3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLdzrLRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jP7EBh2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iup2cbGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xemTioILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZiYoE9lUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RbdbhQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e63IT5uZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3gbT8vfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O64DnnWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpL3vCTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8olMePhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fQzXteYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyx4pRi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcCLxpTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbsESO7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZgjVcL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wWCUrJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q0sV8kTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nN4L9hACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCObl9gKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mf8aejK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UY948uUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEYJqFGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wcwCb6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9q529RulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdYW7wpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zB6lm1HrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zfYteVO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDvZCMCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlZC8dgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnybIlu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASUsCV1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzgqbD0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3Jhx1icWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcxTMakLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMUfM27gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CfCLP5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYL72ilaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7D316TyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJD67vuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rUfG8jBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wgk53O7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEuoIpNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlTGE7FwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1wifhNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzUxOMjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9P89nE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1myBJIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvv1NpaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVZXg2nPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jGgnpWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCy9XSH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFYiKaaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4ckp0pyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lx0yaDlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fivI3hMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teArDXcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEHPtSQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HgpOqcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhKAP1NVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITWZYy10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLmOOgXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AsNj0Dd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sy47nNusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jalf2HSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4KMdBKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRbqCHa5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNJby5sDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ox4vogChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ku3fRX1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxMbJjXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6LzYhn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFUrOIb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6u1Z75WUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DrYgFM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fx7GisBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTQjayJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9lDM9QmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMEHIObfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5v5FLdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2mzCRijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IbiyobMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKUmbz1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFdXQyf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tc3ZLkjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUKcPLA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVLVLu45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aFNrEy3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QANHCBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pR6G1HBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCSOVTtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qa3YDxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZseFfYsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOH1VrhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ru3EHBrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78KV5TdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNyWSGOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eN1qJA2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CWzKNvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRk8MnczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJfi9WZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yb6g3L6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhDF4R1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XceEeIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrSo6SDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08GZ0yDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kaBN14vgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func au5c7ZZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xz65Pt8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k10D4wzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8Ia4RCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gp1367rEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a3yeq3hbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FZLi3VdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgUbrDBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhyN6jBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBGsDU4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2eK2yovzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vrSKAgMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHZAfsbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FT1UpYPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDw2L7pBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAzy6aHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8Go94CtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWf46AJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ryHUty1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrhdyM0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XX3qpGNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hx2u30LAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDMV8HdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func whxDhw2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HpfIbYpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuTmqjLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVJR3MvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwxNfbPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQrPn2GsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HE3sqlmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S481d8wSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySwoefOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41RdxeFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b4qssqGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 10ZPCvg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ryHqbO6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8npqtuuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPUZvMgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQQSwIk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHYJmncfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9ADlRnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpIeWfGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65KmJch7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVIxL4VCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shilwoJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eo6eNgJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HoSNiP4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TmLrkcFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17h6BqbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjsWZVm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGbAXJmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGa4NRKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dOxzbF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXd5XhTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atCBfb2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzxNQGW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BfEL27pJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGdq3OklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KaOOcWdWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nV4nEozCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RC6JkQrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nsUe3K1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jkbiFAxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZI5siRrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCWqca7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94ZYfFvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0HsJi0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxHeLdrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MZeiLe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpAzr4WDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0G7hIpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQUyPnJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1p1DeFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GfAPJlrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZ2kPFKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEJUFE7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRPCcDEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CD3cEukuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJ4RooK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vDXmobMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kG74npzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBh7ockBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hm4Xk7LjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enXHSfnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxuQob8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6J3BAwm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tl9Nvc4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V2cGhhtLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QA64rNkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMVxFZjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNjAXiWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNJLIGQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIdfrTXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icIOpBIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uJHu5IlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5uBzZXCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHMSOC8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKCVjq2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMAtu0KbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4sssScsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JJlk46CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func snUm7kJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8wlklm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebIS9AUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Ch6lQgKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkHDe1WGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZn2uMebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ruSBgVm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7fqDaerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0JEtidbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTiHMch2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OP2S21yBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMkAXsTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68MR8SnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2FDg7VeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ie4fmpopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 334R1dbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X77RC98KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5AdXfxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YCZDzIYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTTpkfuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9hlRwMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8OtwwUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r202r2xkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAiTcAy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUZckq0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxTrxoUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rjnsLk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Utf5O7KDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rlO51sJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDs6RQjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RV0KiIRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQMEcviSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7C5bW9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2NrZqvTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXaprt9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXuEpLdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpYoO6xlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXaGSaOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVRVZrGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9ZGMKxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKmEE9awWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tRM6lVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMMHsMRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEjIAEkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsRhdfnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

