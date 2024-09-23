import { QueryClient } from "react-query";

let queryClientInstance: QueryClient;

export const getQueryClientInstance = () => {
  if (!queryClientInstance)
    queryClientInstance = new QueryClient({
      defaultOptions: {
        queries: {
          staleTime: 5 * 1000,
          retry() {
            return false;
          },
        },
      },
    });

  return queryClientInstance;
};

export const getQueryClient = getQueryClientInstance();
