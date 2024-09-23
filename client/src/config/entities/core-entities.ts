export interface BaseErrorResponse<T extends object = object> {
  message?: string;
  code?: string;
  label?: string;
  data?: T;
  subErrors?: string[];
}

export default class BaseResponse<T = object | void> {
  data?: T;

  message?: string;

  code?: string;

  count?: number;

  limit?: number;

  page?: number;

  constructor(Data?: { new (): T }, code?: string, message?: string) {
    if (code) {
      this.code = code;
    }
    if (message) {
      this.message = message;
    }
    if (Data) {
      this.data = new Data();
    }
  }
}

export interface BaseMetaListResponse {
  currentPage: number;
  itemCount: number;
  itemsPerPage: number;
  totalItems: number;
  totalPages: number;
}

export interface BaseListResponse<T extends object> {
  items?: T;
  meta: BaseMetaListResponse;
}
