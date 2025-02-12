'use client';

import * as React from 'react';
import {
  ClipboardList,
  Command,
  LayoutDashboard,
  Route,
  SatelliteIcon,
  Settings2,
  Shuffle,
  User,
  UserCog,
} from 'lucide-react';

import { NavMain } from '@/components/nav-main';
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

const navMain = [
  {
    title: 'Dashboard',
    url: '/dashboard',
    icon: LayoutDashboard,
    requiredPermissions: ['ViewDashboard'],
  },
  {
    title: 'Registration',
    url: '#',
    icon: ClipboardList,
    requiredPermissions: ['CreateUser', 'CreateVehicle'],
    items: [
      {
        title: 'Car Owner Registration',
        url: '/carOwnerRegistration',
        requiredPermissions: ['CreateUser'],
      },
      {
        title: 'Vehicle Registration',
        url: '/vehicleRegistration',
        requiredPermissions: ['CreateVehicle'],
      },
      {
        title: 'Driver Registration',
        url: '/driver',
        requiredPermissions: ['CreateUser'],
      },
      {
        title: 'Queue Manager Registration',
        url: '/queueManagerForm',
        requiredPermissions: ['CreateUser'],
      },
    ],
  },
  {
    title: 'Users',
    url: '/user',
    icon: User,
    requiredPermissions: ['ViewUser'],
  },
  {
    title: 'Delegation',
    url: '/delegation',
    icon: Shuffle,
    requiredPermissions: ['CreateDelegation'],
  },
  {
    title: 'Route Management',
    url: '#',
    icon: Route,
    requiredPermissions: ['ViewRoute', 'CreateRoute'],
    items: [
      {
        title: 'Stations',
        url: '/stations',
        requiredPermissions: ['ViewRoute'],
      },
      { title: 'Route', url: '/route', requiredPermissions: ['ViewRoute'] },
      { title: 'Paths', url: '/paths', requiredPermissions: ['CreateRoute'] },
      {
        title: 'QueueManagerPaths',
        url: '/QueueManagerPaths',
        requiredPermissions: ['QueueManagerPath'],
      },
    ],
  },
  {
    title: 'User Administration',
    url: '#',
    icon: UserCog,
    requiredPermissions: ['ViewRole'],
    items: [
      { title: 'Users', url: '/user', requiredPermissions: ['ViewUser'] },
      {
        title: 'Role And Permission',
        url: '/rolepermission',
        requiredPermissions: ['ViewRole', 'Assign Permission'],
      },
    ],
  },
  {
    title: 'Real Time Monitoring',
    url: '#',
    icon: SatelliteIcon,
    requiredPermissions: ['VehicleTracking'],
    items: [
      {
        title: 'Track Vehicle',
        url: '/vehicleTracking',
        requiredPermissions: ['VehicleTracking'],
      },
      {
        title: 'Track All Vehicles',
        url: '/trackall',
        requiredPermissions: ['VehicleTracking'],
      },
      {
        title: 'Track Nearby Vehicle',
        url: '/trackNearby',
        requiredPermissions: ['VehicleTracking'],
      },
      {
        title: 'Track Single Vehicle',
        url: '/trackOne',
        requiredPermissions: ['VehicleTracking'],
      },
      {
        title: 'Track Vehicle By Path',
        url: '/trackByPath',
        requiredPermissions: ['VehicleTracking'],
      },
    ],
  },
];

// Function to check if a user has any required permissions
const hasPermission = (
  userPermissions: string[],
  requiredPermissions: string[] = []
) => {
  if (requiredPermissions.length === 0) return true;
  return requiredPermissions.some((perm) => userPermissions.includes(perm));
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const user = useUserStore((state) => state.user);
  const navigate = useNavigate();

  if (!user) {
    navigate('/login');
    return null;
  }

  // Filter menu based on user permissions
  const filteredMenuItems = navMain
    .filter((item) => hasPermission(user.permissions, item.requiredPermissions))
    .map((item) => ({
      ...item,
      items:
        item.items?.filter((subItem) =>
          hasPermission(user.permissions, subItem.requiredPermissions)
        ) || [],
    }));

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
                    Kassech Transportation Management
                  </span>
                  <span className="truncate text-xs">KETAS</span>
                </div>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={filteredMenuItems} />
      </SidebarContent>
      <SidebarFooter>{user ? <NavUser user={user} /> : null}</SidebarFooter>
    </Sidebar>
  );
}
