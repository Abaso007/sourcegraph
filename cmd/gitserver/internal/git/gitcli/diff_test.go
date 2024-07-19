	defaultOpts := git.RawDiffOpts{
		InterHunkContext: 3,
		ContextLines:     3,
	}
		r, err := backend.RawDiff(ctx, "testbase", "HEAD", git.GitDiffComparisonTypeOnlyInHead, defaultOpts)
		r, err := backend.RawDiff(ctx, "testbase", "HEAD", git.GitDiffComparisonTypeIntersection, defaultOpts)
		r, err := backend.RawDiff(ctx, "testbase", "HEAD", git.GitDiffComparisonTypeOnlyInHead, defaultOpts, "f2")
	t.Run("custom context", func(t *testing.T) {
		// Prepare repo state:
		backend := BackendWithRepoCommands(t,
			"echo 'line1\nline2\nline3\nlin4\nline5\nline6\nline7\nline8\n' > f",
			"git add f",
			"git commit -m foo --author='Foo Author <foo@sourcegraph.com>'",
			"git tag testbase",
			"echo 'line1.1\nline2\nline3\nlin4\nline5.5\nline6\nline7\nline8\n' > f",
			"git add f",
			"git commit -m foo --author='Foo Author <foo@sourcegraph.com>'",
		)

		var expectedDiff = []byte(`diff --git f f
index 0ef51c52043997fdd257a0b77d761e9ca58bcc1f..58692a00a73d1f78df00014edf4ef39ef4ba0019 100644
--- f
+++ f
@@ -1 +1 @@
-line1
+line1.1
@@ -5 +5 @@ lin4
-line5
+line5.5
`)

		r, err := backend.RawDiff(ctx, "testbase", "HEAD", git.GitDiffComparisonTypeOnlyInHead, git.RawDiffOpts{
			InterHunkContext: 0,
			ContextLines:     0,
		})
		require.NoError(t, err)
		diff, err := io.ReadAll(r)
		require.NoError(t, err)
		require.NoError(t, r.Close())
		t.Log(string(diff))
		require.Equal(t, string(expectedDiff), string(diff))
	})
		_, err := backend.RawDiff(ctx, "unknown", "test", git.GitDiffComparisonTypeOnlyInHead, defaultOpts)
		_, err = backend.RawDiff(ctx, "test", "unknown", git.GitDiffComparisonTypeOnlyInHead, defaultOpts)
		r, err := backend.RawDiff(ctx, "testbase", "HEAD", git.GitDiffComparisonTypeOnlyInHead, defaultOpts, "/dev/null", "/etc/hosts")
		r, err := backend.RawDiff(ctx, "testbase", "HEAD", git.GitDiffComparisonTypeOnlyInHead, defaultOpts)