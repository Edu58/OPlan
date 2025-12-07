"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";

const Navbar = () => {
  const pathName = usePathname();

  return (
    <nav className="fixed w-full top-0 z-50 blur-backdrop border-b border-gray-200">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          <div className="flex items-center space-x-8">
            <div className="shrink-0">
              <h1 className="text-2xl font-bold text-primary">Oplan</h1>
            </div>
            <div className="hidden md:flex items-center space-x-6">
              <Link
                href="/"
                className={`transition-colors
                  ${pathName == "/" ? "text-primary border-b-2 border-primary" : "text-gray-600 hover:text-gray-900"}`}
              >
                Events
              </Link>
              <Link
                href="/about-us"
                className={`px-3 py-2 text-sm font-medium
                  ${pathName == "/about-us" ? "text-primary border-b-2 border-primary" : "text-gray-600 hover:text-gray-900"}`}
              >
                About
              </Link>
              <Link
                href="/contact-us"
                className={`${pathName == "/contact-us" ? "text-primary border-b-2 border-primary" : "text-gray-600 hover:text-gray-900 transition-colors"}`}
              >
                Contact
              </Link>
            </div>
          </div>
          <div className="flex items-center space-x-4">
            {pathName == "/auth" ? (
              <Link
                href={"/"}
                className="bg-primary text-white px-6 py-2 rounded-full hover:bg-indigo-700 transition-colors font-medium"
              >
                Get Started
              </Link>
            ) : (
              <Link
                href={"/auth"}
                className="text-gray-600 hover:text-gray-900 transition-colors"
              >
                Sign In
              </Link>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
