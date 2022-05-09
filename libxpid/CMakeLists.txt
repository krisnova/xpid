# =========================================================================== #
#            MIT License Copyright (c) 2022 Kris Nóva <kris@nivenly.com>      #
#                                                                             #
#                 ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓                 #
#                 ┃   ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗   ┃                 #
#                 ┃   ████╗  ██║██╔═████╗██║   ██║██╔══██╗  ┃                 #
#                 ┃   ██╔██╗ ██║██║██╔██║██║   ██║███████║  ┃                 #
#                 ┃   ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║  ┃                 #
#                 ┃   ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║  ┃                 #
#                 ┃   ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝  ┃                 #
#                 ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛                 #
#                                                                             #
#                        This machine kills fascists.                         #
#                                                                             #
# =========================================================================== #


######################################################################
set(SO_VERSION "1.2.1")
######################################################################

cmake_minimum_required(VERSION 3.9)
project(xpid VERSION ${SO_VERSION} DESCRIPTION "Linux Process Discovery")

#
# Define all internal files
#
add_library(xpid SHARED
        src/proc.c
        src/proc_dir.c
        src/bpf.c
        )

#find_package(OpenSSL REQUIRED)
#target_link_libraries(xpid OpenSSL::SSL)

#
# Version of the library properties
#
set_target_properties(xpid PROPERTIES VERSION ${SO_VERSION})
#
# This is where the magic symlinking in /usr/include comes to play
set_target_properties(xpid PROPERTIES SOVERSION ${SO_VERSION})

set_target_properties(xpid PROPERTIES SOVERSION ${SO_VERSION})


#
# Everything in /include will be what we expose as API
#
set_target_properties(xpid PROPERTIES PUBLIC_HEADER include/xpid.h)

# Add the top level directory to make our #include statements cleaner
target_include_directories(xpid PRIVATE .)
# Add include (public API)
target_include_directories(xpid PRIVATE include)
# Add src (private implementation)
target_include_directories(xpid PRIVATE src)

# make install
# Here is the install rule
include(GNUInstallDirs)
# This is where we put the shared object on the filesystem
install(TARGETS xpid
        LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
        PUBLIC_HEADER DESTINATION ${CMAKE_INSTALL_INCLUDEDIR})