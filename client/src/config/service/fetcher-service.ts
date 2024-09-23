"use client";

import axios, { AxiosError, AxiosRequestConfig } from "axios";
import qs from "qs";
import { BaseErrorResponse } from "../entities/core-entities";

const coreFetcher = axios.create({
  baseURL: process.env.REACT_APP_BASE_URL ?? "",
  paramsSerializer: (params) => qs.stringfy(params || {}),
});

const interceptors = {
  request: {
    config: async (config: AxiosRequestConfig) => {
      // const token = localStorage && localStorage.getItem("token");

      // if (token && config.headers) {
      //   config.headers.Authorization = `Bearer ${token}`;
      // }
      return config;
    },
  },
  response: {
    error: (error: AxiosError<BaseErrorResponse>) => {
      if (error.response?.status === 401) {
        //   Toast.show(labels.sessionExpire);
        //   RootNavigation.navigate("Auths", {});

        localStorage.clear();
      }

      const response = error.response?.data;

      return Promise.reject(response ?? error);
    },
  },
};

coreFetcher.interceptors.request.use(interceptors.request.config as any);
coreFetcher.interceptors.response.use((response) => {
  const res = response.data as BaseErrorResponse;

  return response;
}, interceptors.response.error);

const fetcher = {
  coreFetcher,
};

export default fetcher;
