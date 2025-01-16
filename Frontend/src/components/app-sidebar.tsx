'use client';

import * as React from 'react';
import {
  BookOpen,
  Bot,
  Command,
  Frame,
  Map,
  PieChart,
  Settings2,
  SquareTerminal,
} from 'lucide-react';

import { NavMain } from '@/components/nav-main';
import { NavProjects } from '@/components/nav-projects';
import { NavUser } from '@/components/nav-user';
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@/components/ui/sidebar';
import { useUserStore } from '@/store/userStore';
import { useNavigate } from 'react-router-dom';

const data = {
  navMain: [
    {
      title: 'Registration',
      url: '#',
      icon: SquareTerminal,
      isActive: true,
      items: [
        {
          title: 'Car Owner Registration',
          url: '/carOwnerRegistration',
        },
        {
          title: 'Vehicle Registration',
          url: '/vehicleRegistration',
        },
        {
          title: 'Driver Registration',
          url: '/driver',
        },
        {
          title: 'Queue Manager Registration',
          url: '/queueManagerForm',
        },
      ],
    },
    {
      title: 'Users',
      url: '#',
      icon: Bot,
      items: [
        {
          title: 'Users',
          url: '/user',
        },
        {
          title: 'Stations',
          url: '/stations',
        },
        {
          title: 'Route',
          url: '/route',
        },
        {
          title: 'Role And Permission',
          url: '/rolepermission',
        },
      ],
    },
    {
      title: 'View Report',
      url: '#',
      icon: BookOpen,
      items: [
        {
          title: 'Introduction',
          url: '#',
        },
        {
          title: 'Get Started',
          url: '#',
        },
        {
          title: 'Tutorials',
          url: '#',
        },
        {
          title: 'Changelog',
          url: '#',
        },
      ],
    },
    {
      title: 'Settings',
      url: '#',
      icon: Settings2,
      items: [
        {
          title: 'General',
          url: '#',
        },
        {
          title: 'Team',
          url: '#',
        },
        {
          title: 'Billing',
          url: '#',
        },
        {
          title: 'Limits',
          url: '#',
        },
      ],
    },
  ],
  //   navSecondary: [
  //     {
  //       title: 'Support',
  //       url: '#',
  //       icon: LifeBuoy,
  //     },
  //     {
  //       title: 'Feedback',
  //       url: '#',
  //       icon: Send,
  //     },
  //   ],
  projects: [
    {
      name: 'Delegation',
      url: '#',
      icon: Frame,
    },
    {
      name: 'Active Driver',
      url: '#',
      icon: PieChart,
    },
    {
      name: 'Report',
      url: '#',
      icon: Map,
    },
  ],
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const user = useUserStore((state) => state.user);
  console.log('🚀 ~ AppSidebar ~ user:', user);
  console.log('🚀 ~ AppSidebar ~ user:', user);
  const navigate = useNavigate();

  if (!user) {
    navigate('/login');
    return;
  }
  return (
    <Sidebar variant="inset" {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild>
              <a href="#">
                <div className="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                  <Command className="size-4" />
                </div>
                <div className="grid flex-1 text-left text-sm leading-tight">
                  <span className="truncate font-semibold">
                    Kassech Transportaion Managment
                  </span>
                  <span className="truncate text-xs">Role</span>
                </div>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
        <NavProjects projects={data.projects} />
        {/* <NavSecondary items={data.navSecondary} className="mt-auto" /> */}
      </SidebarContent>
      <SidebarFooter>{user ? <NavUser user={user} /> : null}</SidebarFooter>
    </Sidebar>
  );
}
