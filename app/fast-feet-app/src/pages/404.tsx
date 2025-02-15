import notFound from "@/assets/human-lost-404.png";

import { Link } from "react-router-dom";

export function NotFound() {
  return (
    <div className="bg-indigo-900 relative overflow-hidden h-screen">
      <img src={notFound} className="absolute h-full w-full object-cover" />
      <div className="inset-0 bg-black opacity-25 absolute" />
      <div className="container mx-auto px-6 md:px-12 relative z-10 flex flex-col items-center justify-center h-full text-center">
        <h1 className="font-extrabold text-5xl text-white leading-tight mt-4">
          Você está sozinho aqui
        </h1>
        <p className="font-extrabold text-8xl my-12 text-white animate-pulse">
          404
        </p>
        <Link
          to="/orders"
          className="bg-white text-indigo-900 font-bold py-3 px-6 rounded-lg shadow-md hover:bg-gray-200 transition"
        >
          Voltar para encomendas
        </Link>
      </div>
    </div>
  );
}
