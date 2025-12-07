import { Lexend } from "next/font/google";
import "./globals.css";
import { Navbar } from "@/components";

const lexend = Lexend({
  subsets: ["latin"],
  variable: "--font-lexend",
  display: "swap",
});

export const metadata = {
  title: "Oplan",
  description: "Connect with your community through extraordinary experiences",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en" className={lexend.variable}>
      <body className={`${lexend.variable} antialiased bg-gray-500`}>
        <Navbar />
        {children}
      </body>
    </html>
  );
}
