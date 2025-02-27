import { CHANGELOG_DIRECTORY } from "@fern-api/configuration-loader";
import { AbsoluteFilePath, doesPathExist, join, RelativeFilePath } from "@fern-api/fs-utils";
import { listFernFiles } from "./listFernFiles";
import { APIChangelog } from "@fern-api/api-workspace-commons";

export async function loadAPIChangelog({
    absolutePathToWorkspace
}: {
    absolutePathToWorkspace: AbsoluteFilePath;
}): Promise<Promise<APIChangelog | undefined>> {
    const absolutePathToChangelogDirectory = join(absolutePathToWorkspace, RelativeFilePath.of(CHANGELOG_DIRECTORY));
    const changelogDirectoryExists = await doesPathExist(absolutePathToChangelogDirectory);
    if (!changelogDirectoryExists) {
        return undefined;
    }

    const mdFiles = await listFernFiles(absolutePathToChangelogDirectory, "{md,mdx}");
    return {
        files: await Promise.all(
            mdFiles.map((file) => {
                return {
                    absoluteFilepath: file.absoluteFilepath,
                    contents: file.fileContents
                };
            })
        )
    };
}
