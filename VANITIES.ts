// Configuration for go.nc0.fr
//
// The prefix `/` is reserved as an index for all the registered modules on the
// server.
// Documentation for `file` and `dir` (substitutions) can be found here:
// https://github.com/golang/gddo/wiki/Source-Code-Links

import { Module } from "./src/index";

/**
 * Cache maximum age, 86400 seconds = 24 hours.
 * The value must be a time, in seconds, we should be respected
 * by clients on requests.
 */
export const MAX_AGE = 86_400;

/**
 * Go modules that will get a vanity URL at go.nc0.fr
 */
export const MODULES: Module[] = [
  /** Example:
   * {
   *  dir: "https://github.com/n1c00o/foo{/dir}",
   *  file: "https://github.com/n1c00o/foo{/dir}/{file}#L{line}",
   *  prefix: "foo",
   *  repo: "https://github.com/n1c00o/foo.git",
   *  svn: "git"
   * }
   */
];
