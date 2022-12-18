import { MODULES, MAX_AGE } from "../VANITIES";

/**
 * Module represents a module's configuration for a vanity URL.
 */
export interface Module {
	/**
	 * The desired module's prefix (import name) without hostname.
	 * Example: foo, to be imported as go.nc0.fr/foo
	 */
	prefix: string;

	/**
	 * The module's source repository URI.
	 * Example: https://github.com/foo/bar.git
	 */
	repo: string;

	/**
	 * The SVN tool used.
	 * Go supports Git, Bazaar, Subversion and Mercurial.
	 */
	svn: "bzr" | "git" | "svn" | "hg";

	/**
	 * A URI to a list of the files in a module's directory.
	 * Example for GitHub: https://github.com/foo/bar/tree/master{/dir}
	 */
	dir: string;

	/**
	 * A URI to the content of a file in a module (even to a line!).
	 * Example for GitHub: https://github.com/foo/bar/tree/master{/dir}/{file}#L{line}
	 */
	file: string;
}

/**
 * Generates the HTML page for the "/" route.
 * @param mds - The list of registered Go modules.
 */
function computeIndexTemplate(mds: Module[]): string {
	const htmlModules: string = mds
		.map(
			(m) =>
				`<li><a href="https://pkg.go.dev/go.nc0.fr/${m.prefix}">go.nc0.fr/${m.prefix}</a></li>`
		)
		.join("\n");

	return `<!DOCTYPE html>
<html>
  <body>
    <h1>Go modules</h1>
    <ul>
      ${htmlModules}
    </ul>
  </body>
</html>`;
}

/**
 * Generates the HTML page for the a module route.
 * @param md - The Go module.
 */
function computeModuleTemplate(md: Module): string {
	return `<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <meta name="go-import" content="go.nc0.fr/${md.prefix} ${md.svn} ${md.repo}">
    <meta name="go-source" content="go.nc0.fr/${md.prefix} ${md.repo} ${md.dir} ${md.file}">
    <meta http-equiv="refresh" content="0; url=https://pkg.go.dev/go.nc0.fr/${md.prefix}"> 
  </head>
  <body>
    <p>Redirecting to <a href="https://pkg.go.dev/go.nc0.fr/${md.prefix}">pkg.go.dev</a></p>
  </body>
</html>`;
}

export default {
	async fetch(req: Request): Promise<Response> {
		const pathname = new URL(req.url).pathname;

		if (pathname === "/") {
			// Index page, a list of all registered modules.
			return new Response(computeIndexTemplate(MODULES), {
				status: 200,
				headers: {
					"Content-Type": "text/html;charset=UTF-8",
					"Cache-Control": `public,max-age=${MAX_AGE}`,
				},
			});
		}
		// A module page.
		const mds = MODULES.filter((v) => v.prefix === `/${pathname}`);
		if (mds.length === 0)
			return new Response(null, {
				status: 404,
				headers: {
					"Content-Type": "text/html;charset=UTF-8",
					"Cache-Control": `public,max-age=${MAX_AGE}`,
				},
			});

		return new Response(computeModuleTemplate(mds[0]), {
			status: 200,
			headers: {
				"Content-Type": "text/html;charset=UTF-8",
				"Cache-Control": `public,max-age=${MAX_AGE}`,
			},
		});
	},
};
