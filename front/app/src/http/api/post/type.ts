export interface Post<T> {
  p: T
  n: T
}

export type PostResponseData = Post<{ p: number, n: number }>
