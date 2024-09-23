import { AxiosInstance } from "axios";

export default abstract class BaseService {
  abstract readonly QUERY_KEY: Record<string, unknown>;
  abstract readonly fetcher: AxiosInstance;
}
