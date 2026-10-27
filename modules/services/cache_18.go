package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.1" )

func WMC6DGQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwKm5hFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fnM1MHeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wWPXk3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmtOFerxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZSzbTaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Stgo4mmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IvK5U8LEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6lZlDU8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FN5XCdUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9oMxLqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VArqa7R6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHvvHEM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIz1A9tOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9XG5pUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 534hZeXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WM0FwZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ib3sHP6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndOu7G3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UmS0MRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7elEYKNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1m4xM0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlWT3N6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eNQJNZTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41VLcOMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANjuf36RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hafTKRnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NMk5pHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0esAUTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIIlYOgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0ee1OMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxIIM69rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqRQ6WO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtQMCTluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vd5GlbtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VeBJ0m5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFu82yZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJeiwK5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYoAMxvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqNcF2f1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hd3zgUcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYVyFzZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZS6vlqhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eWPp1y4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysQOo7RcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZW7VsMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ib0VvneLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7QAdY6g9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjIRlbJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzqu0jewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PaSdvNYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 188d7WbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Izgq09xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hc0ngetMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8d57z2c0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqznWWFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmgR58RBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zC8Bcmi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tiZRgfRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rAkxiePyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRoninX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5TkRlkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PRikiux2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtkR8OJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uewsBGTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qq5Eq8EZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrJLlHqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6NNqJKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKTrxP30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egq3g628Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwvxAJFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mI0zLP8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuSbjrBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQxFXWY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAX8Ox3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEYEVuIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yO6H0lBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZkdul5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKj5sQ7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRRunTHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oz5y7wVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7NG35bOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tMZtVgtaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uys3XsbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UydZ2vSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLzTdzsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJIb72KzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4MbPTEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Isid6vz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BntV0VnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpWKt7eUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IRyn5riWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8ESPCD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BN5lk25nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jS92sY4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUtDNtyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCuovHKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c6ZcEnQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gy2i9f4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcaxC9aQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWeXTgOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1b1P1xdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElUyzyBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGiyyMOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwaOFihqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r63FylqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6Nede0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cs4zdI2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzl6q8opWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2O8JsCpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pATdJOVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func db2rm4ajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVb80JLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDxnMcnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YnW2XvLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0GemaP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0HTipvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6G5YcC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qtxQic3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pm9BxIHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fLbWhDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NXXh0baFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZDpcQvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMqXPm6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVK83yYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IlUWBZhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99Y2SdnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHCygrc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y7GcxMkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrJxBxSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1BdarOrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I6MWhpNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJ9QyUiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sQgLRq5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZsRKgviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xV2crT3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0R6iZfgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7o6ikl0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hj7RM3LlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKwkpNVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZLn4qsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8dEMDLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cVUPT2rUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFeiH2G1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiEIxgvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kblSobd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWeH0kCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNdSeJ2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2cP24Md5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbe9IDZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GzZ8zfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wquAZ7KWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxcHUecWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdgR34n9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMh7HixxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfrf3h4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SazvI5gmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2I51gmuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGEBdrJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQYRTqa2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 673QhiYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVUpBc1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9izHbMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2GWoV6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiWklL5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGAu16PpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZTjbItqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbmqYbaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ryz9bA0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47DbmCP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5xSUZGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iClmBbNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZmNft5U0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sp5t41plWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWTrpx7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDLK0i55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usnv69L4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eALl7SuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GlvMkJCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dvm15sGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CK4N2XndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t03VT2mVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jiseQzWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVffdQLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZkTdCFjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiC91rCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2t1b0SSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxIQj5stWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VM7RCw6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odXYKtyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVbscaZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spglfqwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kDnSymFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PtjoVW0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yu0Ib0F0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0i56kFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8eB5zAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4skn0Wc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGhCoXOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOB8sC2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6XbqeP0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8ZZVErfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SqE4sfMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsFwK5sBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qu55sA9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8irQkiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8TjjB1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CUHelW8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNE8ewZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ege7WNoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ci2k3PlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQqE2GfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHbJQDRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8KHhV50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LApKnuLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZD0lOWihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pz3TUtjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLq230o8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0onqbd8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mf8hwcVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMCrKV2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iacuy1ANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oToLxitMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqSLMg5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fKm1ft5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2DqiwDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6inl64aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HB2yxkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S1rUfdeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgEwYppbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGxlikmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GSmcxh54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3y2anFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNsftbqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPdqWlhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OyJTn8WMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZokZJlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYH87m4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwUc2daJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tIKpojN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ne9fK8zYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPvkqCeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xgk3eh9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3u3dcwncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func of6YhjigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0rjVG9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYe8JvEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3H4bt5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nq4ffbweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khjBMnhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQ9rk9B4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6kCmJNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANcFukuBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tUEcwZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fmtKYZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNTexCIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4olcYXLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVzegQqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dj24Q1F2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VXcbUeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCiieVc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjmxEU9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gW2ZpISgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFRGKbt5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eDojK9RmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ibEFR6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZdSeDkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzRagNm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtVY0VThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAPN4yWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1b0MarIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZV8znf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGRnFtuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5b23BC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ak1QdlV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7APrKDpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnCbRpXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lSEIF8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfs3q9q9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uj5OIPL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9BguPV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ak87t1Q8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elwPS57KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SmvN7nK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MD9VFQGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4vZVQ8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbGDu84FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAGD0u8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaWXbmyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9Usl3oXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfmiZpRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfM2dO3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4SDfFJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQVquneeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cCMQV0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqHXYwxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yyqpmjw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwrLaMeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HnWoNbGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ytm7qSwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKDWcCYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hTzfRq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vk794JDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilB5NwYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ef88urNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NI67IK8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OLhw43JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJK087sIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYtMA4kaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggxahrBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Oz2CgyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LalavLpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkWiauLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Jgyhhl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIMbhk7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NeYkEN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7uQNvM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEZ9DEP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQ1q4ww0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CietI1beWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pn14r2TtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func liz2L2WUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mshf1ePdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djkGnPw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMSaxi4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SmNMgJAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqTWY46uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dt74Iq4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dj5fxQTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mu1yRpsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Go0VlNgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 088O6NlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ms94HGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bhMShvrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tKac7NFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPsNDObiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtQCqljgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhLmmD5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGR3YB0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejv7OHfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65hXhKtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdXodk7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vwNA7iEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6g2c8X0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1hQzeZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGrpjGhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6Ve27E6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUY46svjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9drMLXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LGzuIDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 25e4tkoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8h50BBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7vV3XtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtEBLMnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkdlfwvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuXFp25mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFq1kQMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aF2L7gJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93UkT6UVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4p8Yn8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZwaM9E0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgquPVctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAYyMX2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XB7i5VpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CIlg5zisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhCggloJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zN8WpDk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxpcmRp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZbJyPgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9PzGMEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEGFJ3n5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5YSG9a7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBKEI77IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HSJiS7yUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3E40EfgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blgP1Y18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KohKAwbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XXLPNTRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktgHtf4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZy0fCJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r11pgo6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OA9rytIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1aYRCy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqLm39OhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vu7jbiaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GENroNW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qoeISoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIjNDr7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func we2FsGxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRvlUXMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtAjRt7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwmLVUieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b4BmrzH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqMxuvBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pF9yLLePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9v5nEL0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLbBtMUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boPl3GSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCKYDxeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thW9imr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IfnHNw5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQWK1Kz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xoEyGxS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlBNkp2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2O9fO0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEMOXNWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mk9RWbamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xOMZ6xRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfdAssxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RThjJrP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xp1ZfZEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSWWTo97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzHTywTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haB0p1LkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTQmmtDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufbAFqqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MZU9ljWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lRIF5YxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WshedgHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jpecZDIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZTlwwfGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BK6FZKPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1ojPaXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uJI4SoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8JiPQtN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBwOY9NQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJBaiAdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qc9yLvZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyCZHT6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSMaLv4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9qao7AUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gMGKcJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmK9WVVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bJoTt95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func plKDzDvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzT3oSatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiUBR2PfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACXwLeihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DXiHp4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1U6hFORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cYehDQksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jPtXJXIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fz4FIg3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoepbmOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NL7rv5DQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EPaORKt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iq99uN8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOhotyjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkFqYJMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lno90ooTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFZWjqliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGdotgwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6IL6CqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQ7jEtOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SernJ1R3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func azn3aBCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1aQgkxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuXwoLigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kups5BjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jM9rItliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUhOYlgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dz4P94uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jpqQwE7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRXR3fEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txgvANjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwy3facrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elkyLnazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p6Uo1mmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0emypvYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKhJzu8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6YK3uNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ix3E8D98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0AC1TUVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2OX0YbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHvmuMUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMBATkibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ELTB48Z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wWITdDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMPDBR8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q49Ler5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ye3dWl8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dDKYJLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rx8FrqDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFfw1yHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qQifYPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdLGrRLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJCCSAVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7OKDy9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNquz7UTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzxVXZt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2ZYybsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzoeGyRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 14mE9xP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXn4ZSr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7ZQ0olAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZfYEOKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HYu250t0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvhJuZlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RuI2rqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9kRPfslfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivxtFTaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O0UXrEk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oL5S9GpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZaK0yipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGuyCteLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l67HT7RAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1KJVdFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lROr601Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tUSKSMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9pU78AHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjsFH23jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUR6kESFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7ttPslYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqkNQ9vxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkAKfB2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D9C8QNxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLu8JiGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01kW6yorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iaf9AdWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQkvEKDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2roXnV4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXsax7MPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3igCf0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acoyDYRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBlAiUZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3gHHOKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JeYxeGRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xS0NyYPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eROWhGq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xU1MO7hnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0V5yQ4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cySBCBBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OuxGZVUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BrW6YL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJSZqbVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjhIOH5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svVUVRGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eZqaW8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5eSb5QVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frcIv88qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2ae45xWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzNoVmxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKaG21FLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uYfgD9nLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MjaQyvpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mz5Xw6dDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4uMiqS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uHTaRGIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvtmZaG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVRegGcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OgLCrWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MrnI3G5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZhkno8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rd4OcwjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKsjZS6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyvyLDKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVuzLHp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sb8drsJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUVXlkB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WvINivqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cam589bQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViCno1toWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kpS8DwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s18D83ZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func okgW4Df8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XByEA1y8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxlepUbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0WkFVjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRfqyxmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLcSLnXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wWBOZuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vWtWCMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vH3K4ay6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvNEeesvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnJjV4lzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLLaSxIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjCp7lbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3QsKWVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HX4Ix7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WuuUCuBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHop37eaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSjj3e2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUkwYtw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKySHRwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5fc7wckcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQbNsBeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AecRVXbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfOuKYG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPZLq0f3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0htFSQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9hbkkGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiA3uOkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xf5lf4w2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q67UgTtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMZNYMRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62YvhP6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxJJAPj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uveESPRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Otf9pUOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71K9iIpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R49MzRWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z6ESQOQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPL7AvwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HjK09CS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXjdjgm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfC8VHUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKFeZ3smWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvmEjFmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KHaSjuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7C1sGuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96h2QpgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzNchbeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UqPa6XZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yX4iN7NsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvfNLM6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3oNstJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZxpukcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKbC6EBVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVJwK10YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkhJY5mGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0c4M6ASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgp2dz2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Zb61W5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6fF7qzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HL8Yy8ktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uva5bXWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1K5eWTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDyfusH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S1ZZoGspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nk0KDVF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96yaHIMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kp3ZKprfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9e2RtGnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xy78Lcz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0KwlQeDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UhgxCHnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LS2MeyYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VeZtQu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m705gZ2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUjNNgqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EBCFUqcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WbOYj3f3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJ4kPokGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j2Y3TmZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZkmEo2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1F6oqbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0LzJtBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIZ4Wqo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdQwtUhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmwuVUOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoIC9N4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cUK2Fi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRSvmEt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2KZzXQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkDHbyfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpXJEYjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovWQRrXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYELRBQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVkflwIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3exKCfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2vapcEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZux9rHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Be6J3xPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1scKogdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KVCg5aYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVh9GZd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2rRt6EyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PO8Z8lXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1Vx7wzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4miEpgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUABF3FAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4l6wDVnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y1kE7LobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ikhykIkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCOJaEYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gzQix1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGLarSXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzvgUatQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ry7OTe1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmPXFeE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Z0qPM07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e18T6J8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y24VgzGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ry4dYOaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYJ15GmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CMC68gHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9LDmJvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JuhFLLm8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DspnWnoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EQHY3UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YwwJ9jUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLPlBqJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3z2n3fCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uztNadqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCdwmg9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bp7Ihm8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dK1OrZRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IocAidUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n00i8q4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDmLc6hQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czkpVgDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZkVPZyuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1FDLQlc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esn9fgSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwuFq3GEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8tL9fi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzYG6zUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAnIEA5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08xNxop3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OpM3Nm87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4j72rHchWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func akJzQk9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SlzYyH8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtQdblgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WbQdKp3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbJZNvZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wx122C9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0g3o7TKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJMV6BNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SS5hINhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSFm5Gm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pCyH9a6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IY8wTPHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E07D3tK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KVIMcTwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0eWUGyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1np8QiSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2dQewcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5dHni67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUm6MAWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAHXX9isWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88tVz5TzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ctl6R6CrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLpWbF4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTWHYYvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qgcl3FjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CltQw2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NE8ogkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boqJD3qIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tdxyaKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfA05GumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LuRXrHZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVTLRLveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6aq3ei72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GPTVn2GxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UgvUm9mnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AlzA4GRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCWbnirbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8saeBxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z60pQeWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmrLdS3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4kt95Q9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oCK9pxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIlDNqfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMIyak1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06p9xnAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXvnxDW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54UAkThLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tt6DfQiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBkGIV5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nr0tEpABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpCNmghAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eq6iAplPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYJwF5SjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41ziR7ZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGqmYQm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4JhPlYyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ak6O8779Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TolJ7FppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TajrKtzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sv6mYvL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNxulqmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzKJ5CYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWe6hnsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TsfaN7RcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qchGfAMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qsUoi7XiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxxmEx2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCCkJEcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GuN6r910Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdOpxWfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7B8hojvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fS7SrD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXyxCEQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bB0bzwimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUT6BvM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBudz1uYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jgrmND9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04WQWv2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lzfsX1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZhpoiILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FouYMkPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfEL2gMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGejGAsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayRFU0CFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCgst4KBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilb8Pu8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jfc8ztQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cR2lHo83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBFtfunzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhgnL7RmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjDuy9jaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xmnvev2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVy6S3ywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izbfNEdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jaayAoGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RybNYdoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5UAaKADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FR32yNBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEXECcctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tcL73SW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFpPa5zjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWNnOyvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UJiDB5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDdyixD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ce0bGT5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3lzFN82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PdXXBWZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqUItzhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4afq8J1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func batonQHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzxVhIpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vn0hDvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2snsCiEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h4doSCTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LenPY9CuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lu176RXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIoHbUYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9pQCqNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func liGADRWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfLlGinsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3a55zrAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aF88yo24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lwxeuZtaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4YBPG1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xzs0CM2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFtDLg3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z6petlsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bas3ULY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blKW01ahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oTTQvEmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbSwsbXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3FnnirBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjtU6T2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1cpQz11NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9LfdxTSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ChC6F3sYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCmdfr7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkQtFJrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhAsd0vnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2OD5XRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0i9XJQvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxhYQgjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsWYSLo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Akod1oPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ytZaYGEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tshyYwEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1hIWvl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuQqdhTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 00D4afY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4tRg3sOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CV5n0kWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6etEQkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umaFFVWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEbKhQU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jT7nAqX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NItFtc6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4GtMnMu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLQMB5ZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5uoHqDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27BdWnjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvPoEEKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLK5ghsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDiB02eaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMpoDJaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E65y6QQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YPcXzYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zdvszxQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCxFI5TSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOg3reLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SM5Hau9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzZGOYpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wEdMss6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ay98oMyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IVo4KbAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQW1lwB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z033beBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQh3hLRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tkqt9ME8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXme5hiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEjJ8BGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ngHhVfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCu5f6PcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJREhiX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOFpYE1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jpQA1j3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elkzYXdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJZ0mzsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xp5NiuB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGUpUSxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4VSzJwakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXB1It2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62U5tyJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmY884RkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDfNehOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q61z6mSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSUby8c6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlNI31dRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuXYcDqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiZb2GyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPFdPJKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRYDFmQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHl3krlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ltU1dVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46usp7IGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5K8gGD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5zZwk8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 572ISae7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6cjvKpEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVX9Wa5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMcdId5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6j6jbOIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ue3y5W3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCE7HmfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BH3qnXmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYlXZH53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SVut9USWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZWcYN7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWI4cKfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dr8RRtJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PdPawVBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4q5yU1hNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P635sNGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmTnN6BBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTEvM7ySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3vNIqCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6bTAk5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PCF4lDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DkqyPHfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBUNiApOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YWsb39lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

