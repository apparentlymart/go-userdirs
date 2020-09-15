# go-userdirs

[![godoc reference](https://godoc.org/github.com/apparentlymart/go-userdirs/userdirs?status.svg)](https://godoc.org/github.com/apparentlymart/go-userdirs/userdirs)

This is a small Go library for building suitable paths for storing and locating
application-specific, user-specific configuration files, data files, and
cache files.

It has first-class support for the following operating systems:

- Windows, using the shell's [Known Folders](<https://msdn.microsoft.com/en-us/library/windows/desktop/bb776911(v=vs.85).aspx>) API to locate the application data directories.
- Mac OS X, following the [Library directory layout guidelines](https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileSystemOverview/FileSystemOverview.html#//apple_ref/doc/uid/TP40010672-CH2-SW1).
- Linux, following the [XDG Base Directory specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-0.8.html), currently version 0.8.

The library also supports AIX, DragonFlyBSD, FreeBSD, NetBSD, OpenBSD, and
Solaris by following the XDG specification, treating them like Linux.

---

## Concepts

The library is based on similar principles as the XDG specification, namely:

- There are one or more directories to search for configuration files, one of
  which is preferred and used for creating new configuration files.
- There are one or more directories to search for data files, one of which
  is preferred and used for creating new data files.
- There is exactly one cache directory, which can be used for transient files
  that the application is able to recreate if lost.

The exact locations of these directories will depend on operating system and
system configuration, but we assume that there is at least one directory of
each type. Sometimes the same directory will appear in more than one role.

The distinction between configuration and data files is subtle, but a rule of
thumb is to think about whether it would make sense to keep a particular file
under version control if a user were motivated to keep configuration in a
version control system. Configuration files tend to be relatively small, ideally
in text-based formats that humans could easily read and edit. Data files might
be larger, and may be opaque binary files.

Some configuration and data directories may roam between multiple machines, so
applications should either use host-agnostic file formats there or use other
mechanisms to partition files to avoid collisions, such as using a separate
subdirectory for each supported OS.

The cache directory is commonly (though not guaranteed to be) a local directory
that does not roam between machines. Its contents could be lost at any time,
so applications should store here only files that can be easily recreated when
needed.

## Helpers

The entry-point for the package is
[`userdirs.ForApp`](https://godoc.org/github.com/apparentlymart/go-userdirs/userdirs#ForApp),
which uses some given identifying information to construct application-specific
subdirectory paths under the operating-system-specific base directories.

The [`Dirs`](https://godoc.org/github.com/apparentlymart/go-userdirs/userdirs#Dirs)
object it returns has some helper methods for scanning across the various
directories to find files:

- [`FindConfigFiles`](https://godoc.org/github.com/apparentlymart/go-userdirs/userdirs#Dirs.FindConfigFiles)
  and
  [`FindDataFiles`](https://godoc.org/github.com/apparentlymart/go-userdirs/userdirs#Dirs.FindDataFiles)
  scan over their respective directory lists and find all instances of a
  particular constant sub-path, like a specific configuration file name.
- [`GlobConfigFiles`](https://godoc.org/github.com/apparentlymart/go-userdirs/userdirs#Dirs.GlobConfigFiles)
  and
  [`GlobDataFiles`](https://godoc.org/github.com/apparentlymart/go-userdirs/userdirs#Dirs.GlobDataFiles)
  are similar, but apply a glob pattern to each directory that could potentially
  return multiple results. This could be useful to collect up all files with
  with a particular file extension across all selected directories, for example.

## Contributing

The featureset of this library is considered "done" and so it's unlikely that
any new capabilities will be added. However, bug fixes are much appreciated.

One enhancement area that is open is support for new operating systems. However,
once support for an operating system is added the set of directories returned
for it is largely frozen to ensure that dependent applications can continue
to find files they previously created, so there is a bar for adding a new
operating system: there must be a standards document published by the operating
system vendor that defines the set of rules that this library would then follow
when running on that operating system, with confidence that the standards are
long-lived and not likely to need breaking changes in future.

If you'd like to contribute support for a new operating system, please open
an issue to discuss it. Note also that we currently use Travis-CI for testing,
and so we're constrained for what platforms we can run unit tests on. An
operating system without native Travis-CI support would need to be particularly
compelling for it to be accepted, since we would need to find some other way
to automatically test it.

Any pull requests are assumed to be submitted under the same MIT License as
the library is already shared under, as described in [LICENSE](./LICENSE).
