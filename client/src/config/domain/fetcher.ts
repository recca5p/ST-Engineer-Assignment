export interface FetcherPagination {
  page: number;
  size: number;
  total: number;
}

export interface FetcherResult<T = unknown> {
  data: T;
  message: string;
  code: string;
  pagination?: FetcherPagination;
}

export interface BasePaginationRequest {
  limit?: number;
  page?: number;
}

export type FetcherError = FetcherResult<unknown>;
