import React from "react";
import { Link } from "react-router-dom";

interface CardWrapperProps {
  label: string;
  title: string;
  backButtonHref: string;
  backButtonLabel: string;
  children: React.ReactNode;
}

const CardWrapper: React.FC<CardWrapperProps> = ({
  label,
  title,
  backButtonHref,
  backButtonLabel,
  children,
}) => {
  return (
    <div className="w-full max-w-md mx-auto my-auto bg-white shadow-md rounded-md overflow-hidden">
      <div className="px-6 py-4 border-b">
        <div className="text-center">
          <h2 className="text-xl font-semibold">{title}</h2>
          <p className="text-gray-600 text-sm">{label}</p>
        </div>
      </div>
      <div className="p-6">{children}</div>
      <div className="px-6 py-4 bg-gray-100 text-center">
        <Link
          to={backButtonHref}
          className="text-blue-950 hover:text-blue-800 text-sm font-medium"
        >
          {backButtonLabel}
        </Link>
      </div>
    </div>
  );
};

export default CardWrapper;
