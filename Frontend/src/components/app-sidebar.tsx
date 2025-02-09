'use client';

import * as React from 'react';
import {
  ClipboardList,
  Command,
  Frame,
  Map,
  PieChart,
  Route,
  SatelliteIcon,
  Settings2,
  Shuffle,
  User,
  UserCog,
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
      icon: ClipboardList,
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
      url: '/user',
      icon: User,
    },

    {
      title: 'Delegation',
      url: '/delegation',
      icon: Shuffle,
    },

    {
      title: 'Route Management',
      url: '#',
      icon: Route,
      items: [
        {
          title: 'Stations',
          url: '/stations',
        },
        {
          title: 'Route',
          url: '/route',
        },
        {
          title: 'Paths',
          url: '/paths',
        },
        {
          title: 'QueueManagerPaths',
          url: '/QueueManagerPaths',
        },
      ],
    },

    {
      title: 'User Administration',
      url: '#',
      icon: UserCog,
      items: [
        {
          title: 'Users',
          url: '/user',
        },
        {
          title: 'Role And Permission',
          url: '/rolepermission',
        },
      ],
    },

    {
      title: 'Real Time Monitoring',
      url: '#',
      icon: SatelliteIcon,
      items: [
        {
          title: 'Track Vehicle',
          url: '/vehicleTracking',
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

  // projects: [
  //   {
  //     name: 'Delegation',
  //     url: '#',
  //     icon: Frame,
  //   },
  //   {
  //     name: 'Active Driver',
  //     url: '#',
  //     icon: PieChart,
  //   },
  //   {
  //     name: 'Report',
  //     url: '#',
  //     icon: Map,
  //   },
  // ],
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
                  <span className="truncate text-xs">KETAS</span>
                </div>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
        {/* <NavProjects projects={data.projects} /> */}
        {/* <NavSecondary items={data.navSecondary} className="mt-auto" /> */}
      </SidebarContent>
      <SidebarFooter>{user ? <NavUser user={user} /> : null}</SidebarFooter>
    </Sidebar>
  );
}
