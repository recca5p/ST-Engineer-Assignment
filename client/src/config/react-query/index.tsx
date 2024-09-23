"use client";

import { ReactNode, useState } from "react";
import { QueryClientProvider } from "react-query";

import { getQueryClientInstance } from "./query";

export interface ReactQueryProviderProps {
  children: ReactNode;
}

function ReactQueryProvider({ children }: ReactQueryProviderProps) {
  const [queryClient] = useState(() => getQueryClientInstance());

  return (
    // <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
}

export default ReactQueryProvider;
