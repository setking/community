export interface Community<T> {
  p: T
  n: T
}

export type CommunityResponseData = Community<{ p: number, n: number }>
