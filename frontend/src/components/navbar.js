const Navbar = () => {
  return (
    <nav className="fixed w-full top-0 z-50 blur-backdrop border-b border-gray-200">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          <div className="flex items-center space-x-8">
            <div className="shrink-0">
              <h1 className="text-2xl font-bold text-primary">Oplan</h1>
            </div>
            <div className="hidden md:flex items-center space-x-6">
              <a
                href="/events"
                className="text-gray-600 hover:text-gray-900 transition-colors"
              >
                Events
              </a>
              <a
                href="#"
                className="text-primary border-b-2 border-primary px-3 py-2 text-sm font-medium"
              >
                About
              </a>
              <a
                href="#"
                className="text-gray-600 hover:text-gray-900 transition-colors"
              >
                Contact
              </a>
            </div>
          </div>
          <div className="flex items-center space-x-4">
            <button className="text-gray-600 hover:text-gray-900 transition-colors">
              Sign In
            </button>
            <button className="bg-primary text-white px-6 py-2 rounded-full hover:bg-indigo-700 transition-colors font-medium">
              Get Started
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
