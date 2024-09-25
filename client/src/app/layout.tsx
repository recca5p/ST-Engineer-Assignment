"use client";

import type { Metadata } from "next";
import localFont from "next/font/local";
import "@asseinfo/react-kanban/dist/styles.css";

import "./globals.css";
import ReactQueryProvider from "@/config/react-query";
import { useState } from "react";
import { getQueryClientInstance } from "@/config/react-query/query";
import { QueryClient, QueryClientProvider } from 'react-query';


const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export default function RootLayout({
                                     children,
                                   }: {
  children: React.ReactNode;
}) {
  const queryClient = useState(() => new QueryClient())[0];

  return (
      <html lang="en">
      <body
          className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
      <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
      </body>
      </html>
  );
}