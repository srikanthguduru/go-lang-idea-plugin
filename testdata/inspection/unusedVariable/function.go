package main

import "io"

func reverse(a, b int) (c, d int) {
    c, d = b, a
	return
}

func gen1() (g1v int) {
    g1v = 5
    return
}

func gen2() (g2v int) {
    return 6
}

func noop() {
}

func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrint(a, true, true)
	n64, err := w.Write(p.buf)
	p.free()
	return int(n64), err
}

func coroutineTest(/*begin*/a/*end.Unused parameter 'a'*/ int) int {
    go func(a, /*begin*/b/*end.Unused parameter 'b'*/ int) int {
        return a
    }()
    return 2
}

func main() {
    a, b := reverse(1, 2)
    println(a, b, gen1(), gen2())
}
