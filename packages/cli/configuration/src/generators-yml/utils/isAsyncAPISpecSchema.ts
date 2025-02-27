import { AsyncApiSchema, AsyncApiOrOpenApiSpecSchema } from "../schemas";

export function isAsyncAPISchema(spec: AsyncApiOrOpenApiSpecSchema): spec is AsyncApiSchema {
    return (spec as AsyncApiSchema)?.asyncapi != null;
}
