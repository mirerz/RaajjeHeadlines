import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

const inter = Inter({
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "adhu.space | Fluid Sovereignty",
  description: "High-fidelity latency-zero national news broadcast synthesized by Oivaru.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={`${inter.className} h-full antialiased`}>
      <body>
        {children}
      </body>
    </html>
  );
}
