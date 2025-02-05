'use client';

import * as React from 'react';
import { useLoginLogs } from '@/services/dashboardService';
import {
  ColumnDef,
  SortingState,
  getCoreRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
  flexRender,
} from '@tanstack/react-table';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationPrevious,
  PaginationNext,
} from '@/components/ui/pagination';
import { ArrowUpDown } from 'lucide-react';

const LogTable = () => {
  const { data: logs = [], isLoading, error } = useLoginLogs();

  const columns: ColumnDef<(typeof logs)[0]>[] = [
    {
      accessorKey: 'ID',
      header: 'ID',
    },
    {
      accessorKey: 'LoginTime',
      header: 'Login Time',
      cell: ({ row }) => new Date(row.getValue('LoginTime')).toLocaleString(),
    },
    {
      accessorKey: 'IP',
      header: 'IP Address',
    },
    {
      accessorKey: 'UserAgent',
      header: 'User Agent',
    },
    {
      accessorKey: 'UserID',
      header: 'User ID',
    },
    {
      accessorKey: 'UpdatedAt',
      header: 'Updated At',
    },
    {
      accessorKey: 'CreatedAt',
      header: 'Created At',
    },
    {
      accessorKey: 'DeletedAt',
      header: 'Deleted At',
    },
  ];

  const [sorting, setSorting] = React.useState<SortingState>([]);
  const [pageIndex, setPageIndex] = React.useState(0);

  const table = useReactTable({
    data: logs,
    columns,
    state: {
      sorting,
    },
    onSortingChange: setSorting,
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
  });

  if (isLoading) return <p>Loading logs...</p>;
  if (error) return <p className="text-red-500">Error fetching logs.</p>;

  return (
    <div className="p-4 border rounded-lg shadow-md">
      <h2 className="text-xl font-bold mb-4">Login Logs</h2>
      <Table>
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => (
                <TableHead key={header.id}>
                  {header.isPlaceholder
                    ? null
                    : flexRender(
                        header.column.columnDef.header,
                        header.getContext()
                      )}
                </TableHead>
              ))}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows.length ? (
            table.getRowModel().rows.map((row) => (
              <TableRow key={row.id}>
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                No logs found.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
      <Pagination className="mt-4">
        <PaginationContent>
          <PaginationItem>
            <PaginationPrevious
              onClick={() => table.previousPage()}
              className={
                table.getCanPreviousPage()
                  ? ''
                  : 'pointer-events-none opacity-50'
              }
            >
              Previous
            </PaginationPrevious>
          </PaginationItem>
          <PaginationItem>
            <PaginationLink href="#" onClick={(e) => e.preventDefault()}>
              {pageIndex + 1}
            </PaginationLink>
          </PaginationItem>
          <PaginationItem>
            <PaginationNext
              onClick={() => table.nextPage()}
              className={
                table.getCanNextPage() ? '' : 'pointer-events-none opacity-50'
              }
            >
              Next
            </PaginationNext>
          </PaginationItem>
        </PaginationContent>
      </Pagination>
    </div>
  );
};

export default LogTable;
