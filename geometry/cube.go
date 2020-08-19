package geometry

func CubeVolume(n int) int {
  if n != 0 {
    return n * n * n
  } else {
    return 0, errors.New("Zero length edge is not allowed")
  }
}
