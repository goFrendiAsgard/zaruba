
export function getName(req: any) {
    return req.params.name || req.query.name || req.body.name || "";
}
