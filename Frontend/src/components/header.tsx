// components/ui/Header.tsx
import { Breadcrumb, BreadcrumbItem, BreadcrumbLink, BreadcrumbList, BreadcrumbSeparator } from "@/components/ui/breadcrumb";
import { Separator } from "@/components/ui/separator";
import { SidebarTrigger } from "@/components/ui/sidebar";
import React from "react";

// Type definition for breadcrumb item
type BreadcrumbPath = {
  name: string;
  href: string;
};

// Props type for the Header component
interface HeaderProps {
  paths: BreadcrumbPath[];  // Array of breadcrumb paths
}

const Header: React.FC<HeaderProps> = ({ paths }) => {
  return (
    <header className="flex h-16 shrink-0 items-center gap-2">
      <div className="flex items-center gap-2 px-4">
        <SidebarTrigger className="-ml-1" />
        <Separator orientation="vertical" className="mr-2 h-4" />
        <Breadcrumb>
          <BreadcrumbList>
            {paths.map((path, index) => (
              <React.Fragment key={index}>
                <BreadcrumbItem className="hidden md:block">
                  <BreadcrumbLink href={path.href}>{path.name}</BreadcrumbLink>
                </BreadcrumbItem>
                {/* Add separator except for the last item */}
                {index < paths.length - 1 && (
                  <BreadcrumbSeparator className="hidden md:block" />
                )}
              </React.Fragment>
            ))}
          </BreadcrumbList>
        </Breadcrumb>
      </div>
    </header>
  );
};

export default Header;
