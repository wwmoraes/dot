package tokenizer

const (
	// KeywordStrict denotes a strict graph
	KeywordStrict = "strict"
	// KeywordDigraph denotes a directed root graph
	KeywordDigraph = "digraph"
	// KeywordGraph denotes an undirected root graph
	KeywordGraph = "graph"
	// KeywordSubgraph denotes a subgraph
	KeywordSubgraph = "subgraph"
	// KeywordLineComment denotes single line comment
	KeywordLineComment = "//"
	// KeywordOpenBlockComment denotes the start of a multi-line block comment
	KeywordOpenBlockComment = "/*"
	// KeywordCloseBlockComment denotes the end of a multi-line block comment
	KeywordCloseBlockComment = "*/"
	KeywordOpenBlock         = "{"
	KeywordCloseBlock        = "}"
	KeywordSemicolon         = ";"
	KeywordOpenSquareBlock   = "["
	KeywordCloseSquareBlock  = "]"
	KeywordDirectedEdge      = "->"
	KeywordUndirectedEdge    = "--"
	KeywordEquals            = "="

	KeywordAttributeBackground         = "_background"
	KeywordAttributeArea               = "area"
	KeywordAttributeArrowHead          = "arrowhead"
	KeywordAttributeArrowSize          = "arrowsize"
	KeywordAttributeArrowTail          = "arrowtail"
	KeywordAttributeBb                 = "bb"
	KeywordAttributeBgcolor            = "bgcolor"
	KeywordAttributeCenter             = "center"
	KeywordAttributeCharset            = "charset"
	KeywordAttributeClass              = "class"
	KeywordAttributeClusterRank        = "clusterrank"
	KeywordAttributeColor              = "color"
	KeywordAttributeColorScheme        = "colorscheme"
	KeywordAttributeComment            = "comment"
	KeywordAttributeCompound           = "compound"
	KeywordAttributeConcentrate        = "concentrate"
	KeywordAttributeConstraint         = "constraint"
	KeywordAttributeDamping            = "Damping"
	KeywordAttributeDecorate           = "decorate"
	KeywordAttributeDefaultDist        = "defaultdist"
	KeywordAttributeDim                = "dim"
	KeywordAttributeDimen              = "dimen"
	KeywordAttributeDir                = "dir"
	KeywordAttributeDirEdgeConstraints = "diredgeconstraints"
	KeywordAttributeDistortion         = "distortion"
	KeywordAttributeDpi                = "dpi"
	KeywordAttributeEdgeHref           = "edgehref"
	KeywordAttributeEdgeTarget         = "edgetarget"
	KeywordAttributeEdgeTooltip        = "edgetooltip"
	KeywordAttributeEdgeURL            = "edgeURL"
	KeywordAttributeEpsilon            = "epsilon"
	KeywordAttributeEsep               = "esep"
	KeywordAttributeFillColor          = "fillcolor"
	KeywordAttributeFixedSize          = "fixedsize"
	KeywordAttributeFontColor          = "fontcolor"
	KeywordAttributeFontName           = "fontname"
	KeywordAttributeFontNames          = "fontnames"
	KeywordAttributeFontPath           = "fontpath"
	KeywordAttributeFontSize           = "fontsize"
	KeywordAttributeForceLabels        = "forcelabels"
	KeywordAttributeGradientAngle      = "gradientangle"
	KeywordAttributeGroup              = "group"
	KeywordAttributeHeadLp             = "head_lp"
	KeywordAttributeHeadClip           = "headclip"
	KeywordAttributeHeadHref           = "headhref"
	KeywordAttributeHeadLabel          = "headlabel"
	KeywordAttributeHeadPort           = "headport"
	KeywordAttributeHeadTarget         = "headtarget"
	KeywordAttributeHeadTooltip        = "headtooltip"
	KeywordAttributeHeadURL            = "headURL"
	KeywordAttributeHeight             = "height"
	KeywordAttributeHref               = "href"
	KeywordAttributeID                 = "id"
	KeywordAttributeImage              = "image"
	KeywordAttributeImagePath          = "imagepath"
	KeywordAttributeImagePos           = "imagepos"
	KeywordAttributeImageScale         = "imagescale"
	KeywordAttributeInputScale         = "inputscale"
	KeywordAttributeK                  = "K"
	KeywordAttributeLabel              = "label"
	KeywordAttributeLabelScheme        = "label_scheme"
	KeywordAttributeLabelAngle         = "labelangle"
	KeywordAttributeLabelDistance      = "labeldistance"
	KeywordAttributeLabelFloat         = "labelfloat"
	KeywordAttributeLabelFontColor     = "labelfontcolor"
	KeywordAttributeLabelFontName      = "labelfontname"
	KeywordAttributeLabelFontSize      = "labelfontsize"
	KeywordAttributeLabelHref          = "labelhref"
	KeywordAttributeLabelJust          = "labeljust"
	KeywordAttributeLabelLoc           = "labelloc"
	KeywordAttributeLabelTarget        = "labeltarget"
	KeywordAttributeLabelTooltip       = "labeltooltip"
	KeywordAttributeLabelURL           = "labelURL"
	KeywordAttributeLandscape          = "landscape"
	KeywordAttributeLayer              = "layer"
	KeywordAttributeLayerListSep       = "layerlistsep"
	KeywordAttributeLayers             = "layers"
	KeywordAttributeLayerSelect        = "layerselect"
	KeywordAttributeLayerSep           = "layersep"
	KeywordAttributeLayout             = "layout"
	KeywordAttributeLen                = "len"
	KeywordAttributeLevels             = "levels"
	KeywordAttributeLevelsGap          = "levelsgap"
	KeywordAttributeLHead              = "lhead"
	KeywordAttributeLHeight            = "lheight"
	KeywordAttributeLp                 = "lp"
	KeywordAttributeLTail              = "ltail"
	KeywordAttributeLWidth             = "lwidth"
	KeywordAttributeMargin             = "margin"
	KeywordAttributeMaxIter            = "maxiter"
	KeywordAttributeMcLimit            = "mclimit"
	KeywordAttributeMinDist            = "mindist"
	KeywordAttributeMinLen             = "minlen"
	KeywordAttributeMode               = "mode"
	KeywordAttributeModel              = "model"
	KeywordAttributeMosek              = "mosek"
	KeywordAttributeNewRank            = "newrank"
	KeywordAttributeNodeSep            = "nodesep"
	KeywordAttributeNoJustify          = "nojustify"
	KeywordAttributeNormalize          = "normalize"
	KeywordAttributeNoTranslate        = "notranslate"
	KeywordAttributeNsLimit            = "nslimit"
	KeywordAttributeNsLimit1           = "nslimit1"
	KeywordAttributeOrdering           = "ordering"
	KeywordAttributeOrientation        = "orientation"
	KeywordAttributeOutputOrder        = "outputorder"
	KeywordAttributeOverlap            = "overlap"
	KeywordAttributeOverlapScaling     = "overlap_scaling"
	KeywordAttributeOverlapShrink      = "overlap_shrink"
	KeywordAttributePack               = "pack"
	KeywordAttributePackMode           = "packmode"
	KeywordAttributePad                = "pad"
	KeywordAttributePage               = "page"
	KeywordAttributePageDir            = "pagedir"
	KeywordAttributePenColor           = "pencolor"
	KeywordAttributePenWidth           = "penwidth"
	KeywordAttributePeripheries        = "peripheries"
	KeywordAttributePin                = "pin"
	KeywordAttributePos                = "pos"
	KeywordAttributeQuadTree           = "quadtree"
	KeywordAttributeQuantum            = "quantum"
	KeywordAttributeRank               = "rank"
	KeywordAttributeRankDir            = "rankdir"
	KeywordAttributeRankSep            = "ranksep"
	KeywordAttributeRatio              = "ratio"
	KeywordAttributeRects              = "rects"
	KeywordAttributeRegular            = "regular"
	KeywordAttributeRemincross         = "remincross"
	KeywordAttributeRepulsiveforce     = "repulsiveforce"
	KeywordAttributeResolution         = "resolution"
	KeywordAttributeRoot               = "root"
	KeywordAttributeRotate             = "rotate"
	KeywordAttributeRotation           = "rotation"
	KeywordAttributeSameHead           = "samehead"
	KeywordAttributeSameTail           = "sametail"
	KeywordAttributeSamplePoints       = "samplepoints"
	KeywordAttributeScale              = "scale"
	KeywordAttributeSearchSize         = "searchsize"
	KeywordAttributeSep                = "sep"
	KeywordAttributeShape              = "shape"
	KeywordAttributeShapeFile          = "shapefile"
	KeywordAttributeShowBoxes          = "showboxes"
	KeywordAttributeSides              = "sides"
	KeywordAttributeSize               = "size"
	KeywordAttributeSkew               = "skew"
	KeywordAttributeSmoothing          = "smoothing"
	KeywordAttributeSortv              = "sortv"
	KeywordAttributeSplines            = "splines"
	KeywordAttributeStart              = "start"
	KeywordAttributeStyle              = "style"
	KeywordAttributeStylesheet         = "stylesheet"
	KeywordAttributeTailLp             = "tail_lp"
	KeywordAttributeTailClip           = "tailclip"
	KeywordAttributeTailHref           = "tailhref"
	KeywordAttributeTailLabel          = "taillabel"
	KeywordAttributeTailPort           = "tailport"
	KeywordAttributeTailTarget         = "tailtarget"
	KeywordAttributeTailTooltip        = "tailtooltip"
	KeywordAttributeTailURL            = "tailURL"
	KeywordAttributeTarget             = "target"
	KeywordAttributeTooltip            = "tooltip"
	KeywordAttributeTrueColor          = "truecolor"
	KeywordAttributeURL                = "URL"
	KeywordAttributeVertices           = "vertices"
	KeywordAttributeViewport           = "viewport"
	KeywordAttributeVoroMargin         = "voro_margin"
	KeywordAttributeWeight             = "weight"
	KeywordAttributeWidth              = "width"
	KeywordAttributeXdotVersion        = "xdotversion"
	KeywordAttributeXlabel             = "xlabel"
	KeywordAttributeXlp                = "xlp"
	KeywordAttributeZ                  = "z"
)

var KeywordAttributes = [...]string{
	KeywordAttributeBackground,
	KeywordAttributeArea,
	KeywordAttributeArrowHead,
	KeywordAttributeArrowSize,
	KeywordAttributeArrowTail,
	KeywordAttributeBb,
	KeywordAttributeBgcolor,
	KeywordAttributeCenter,
	KeywordAttributeCharset,
	KeywordAttributeClass,
	KeywordAttributeClusterRank,
	KeywordAttributeColor,
	KeywordAttributeColorScheme,
	KeywordAttributeComment,
	KeywordAttributeCompound,
	KeywordAttributeConcentrate,
	KeywordAttributeConstraint,
	KeywordAttributeDamping,
	KeywordAttributeDecorate,
	KeywordAttributeDefaultDist,
	KeywordAttributeDim,
	KeywordAttributeDimen,
	KeywordAttributeDir,
	KeywordAttributeDirEdgeConstraints,
	KeywordAttributeDistortion,
	KeywordAttributeDpi,
	KeywordAttributeEdgeHref,
	KeywordAttributeEdgeTarget,
	KeywordAttributeEdgeTooltip,
	KeywordAttributeEdgeURL,
	KeywordAttributeEpsilon,
	KeywordAttributeEsep,
	KeywordAttributeFillColor,
	KeywordAttributeFixedSize,
	KeywordAttributeFontColor,
	KeywordAttributeFontName,
	KeywordAttributeFontNames,
	KeywordAttributeFontPath,
	KeywordAttributeFontSize,
	KeywordAttributeForceLabels,
	KeywordAttributeGradientAngle,
	KeywordAttributeGroup,
	KeywordAttributeHeadLp,
	KeywordAttributeHeadClip,
	KeywordAttributeHeadHref,
	KeywordAttributeHeadLabel,
	KeywordAttributeHeadPort,
	KeywordAttributeHeadTarget,
	KeywordAttributeHeadTooltip,
	KeywordAttributeHeadURL,
	KeywordAttributeHeight,
	KeywordAttributeHref,
	KeywordAttributeID,
	KeywordAttributeImage,
	KeywordAttributeImagePath,
	KeywordAttributeImagePos,
	KeywordAttributeImageScale,
	KeywordAttributeInputScale,
	KeywordAttributeK,
	KeywordAttributeLabel,
	KeywordAttributeLabelScheme,
	KeywordAttributeLabelAngle,
	KeywordAttributeLabelDistance,
	KeywordAttributeLabelFloat,
	KeywordAttributeLabelFontColor,
	KeywordAttributeLabelFontName,
	KeywordAttributeLabelFontSize,
	KeywordAttributeLabelHref,
	KeywordAttributeLabelJust,
	KeywordAttributeLabelLoc,
	KeywordAttributeLabelTarget,
	KeywordAttributeLabelTooltip,
	KeywordAttributeLabelURL,
	KeywordAttributeLandscape,
	KeywordAttributeLayer,
	KeywordAttributeLayerListSep,
	KeywordAttributeLayers,
	KeywordAttributeLayerSelect,
	KeywordAttributeLayerSep,
	KeywordAttributeLayout,
	KeywordAttributeLen,
	KeywordAttributeLevels,
	KeywordAttributeLevelsGap,
	KeywordAttributeLHead,
	KeywordAttributeLHeight,
	KeywordAttributeLp,
	KeywordAttributeLTail,
	KeywordAttributeLWidth,
	KeywordAttributeMargin,
	KeywordAttributeMaxIter,
	KeywordAttributeMcLimit,
	KeywordAttributeMinDist,
	KeywordAttributeMinLen,
	KeywordAttributeMode,
	KeywordAttributeModel,
	KeywordAttributeMosek,
	KeywordAttributeNewRank,
	KeywordAttributeNodeSep,
	KeywordAttributeNoJustify,
	KeywordAttributeNormalize,
	KeywordAttributeNoTranslate,
	KeywordAttributeNsLimit,
	KeywordAttributeNsLimit1,
	KeywordAttributeOrdering,
	KeywordAttributeOrientation,
	KeywordAttributeOutputOrder,
	KeywordAttributeOverlap,
	KeywordAttributeOverlapScaling,
	KeywordAttributeOverlapShrink,
	KeywordAttributePack,
	KeywordAttributePackMode,
	KeywordAttributePad,
	KeywordAttributePage,
	KeywordAttributePageDir,
	KeywordAttributePenColor,
	KeywordAttributePenWidth,
	KeywordAttributePeripheries,
	KeywordAttributePin,
	KeywordAttributePos,
	KeywordAttributeQuadTree,
	KeywordAttributeQuantum,
	KeywordAttributeRank,
	KeywordAttributeRankDir,
	KeywordAttributeRankSep,
	KeywordAttributeRatio,
	KeywordAttributeRects,
	KeywordAttributeRegular,
	KeywordAttributeRemincross,
	KeywordAttributeRepulsiveforce,
	KeywordAttributeResolution,
	KeywordAttributeRoot,
	KeywordAttributeRotate,
	KeywordAttributeRotation,
	KeywordAttributeSameHead,
	KeywordAttributeSameTail,
	KeywordAttributeSamplePoints,
	KeywordAttributeScale,
	KeywordAttributeSearchSize,
	KeywordAttributeSep,
	KeywordAttributeShape,
	KeywordAttributeShapeFile,
	KeywordAttributeShowBoxes,
	KeywordAttributeSides,
	KeywordAttributeSize,
	KeywordAttributeSkew,
	KeywordAttributeSmoothing,
	KeywordAttributeSortv,
	KeywordAttributeSplines,
	KeywordAttributeStart,
	KeywordAttributeStyle,
	KeywordAttributeStylesheet,
	KeywordAttributeTailLp,
	KeywordAttributeTailClip,
	KeywordAttributeTailHref,
	KeywordAttributeTailLabel,
	KeywordAttributeTailPort,
	KeywordAttributeTailTarget,
	KeywordAttributeTailTooltip,
	KeywordAttributeTailURL,
	KeywordAttributeTarget,
	KeywordAttributeTooltip,
	KeywordAttributeTrueColor,
	KeywordAttributeURL,
	KeywordAttributeVertices,
	KeywordAttributeViewport,
	KeywordAttributeVoroMargin,
	KeywordAttributeWeight,
	KeywordAttributeWidth,
	KeywordAttributeXdotVersion,
	KeywordAttributeXlabel,
	KeywordAttributeXlp,
	KeywordAttributeZ,
}
