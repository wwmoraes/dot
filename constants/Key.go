package constants

// Key attribute key type
type Key string

const (
	// KeyBackground xdot format string specifying an arbitrary background
	//
	// Valid for: Graphs.
	KeyBackground Key = "_background"
	// KeyArea indicates the preferred area for a node or empty cluster when
	// laid out by patchwork
	//
	// Valid for: Clusters, Nodes. Note: patchwork only
	KeyArea Key = "area"
	// KeyArrowHead style of arrowhead on the head node of an edge. This
	// will only appear if the dir attribute is forward or both
	//
	// Valid for: Edges.
	KeyArrowHead Key = "arrowhead"
	// KeyArrowSize multiplicative scale factor for arrowheads
	//
	// Valid for: Edges.
	KeyArrowSize Key = "arrowsize"
	// KeyArrowTail style of arrowhead on the tail node of an edge.
	// This will only appear if the dir attribute is back or both.
	//
	// Valid for: Edges.
	KeyArrowTail Key = "arrowtail"
	// KeyBb bounding box of drawing in points
	//
	// Valid for: Graphs. Note: write only
	KeyBb Key = "bb"
	// KeyBgcolor When attached to the root graph, this color is used as the
	// background for entire canvas.
	//
	// When a cluster attribute, it is used as the initial background for the
	// cluster. If a cluster has a filled style, the cluster’s fillcolor will
	// overlay the background color.
	//
	// If the value is a colorList, a gradient fill is used. By default, this is a
	// linear fill; setting style=radial will cause a radial fill. Only two colors
	// are used. If the second color (after a colon) is missing, the default color
	// is used for it. See also the gradientangle attribute for setting the
	// gradient angle.
	KeyBgcolor            Key = "bgcolor"
	KeyCenter             Key = "center"
	KeyCharset            Key = "charset"
	KeyClass              Key = "class"
	KeyClusterRank        Key = "clusterrank"
	KeyColor              Key = "color"
	KeyColorScheme        Key = "colorscheme"
	KeyComment            Key = "comment"
	KeyCompound           Key = "compound"
	KeyConcentrate        Key = "concentrate"
	KeyConstraint         Key = "constraint"
	KeyDamping            Key = "Damping"
	KeyDecorate           Key = "decorate"
	KeyDefaultDist        Key = "defaultdist"
	KeyDim                Key = "dim"
	KeyDimen              Key = "dimen"
	KeyDir                Key = "dir"
	KeyDirEdgeConstraints Key = "diredgeconstraints"
	KeyDistortion         Key = "distortion"
	KeyDpi                Key = "dpi"
	KeyEdgeHref           Key = "edgehref"
	KeyEdgeTarget         Key = "edgetarget"
	KeyEdgeTooltip        Key = "edgetooltip"
	KeyEdgeURL            Key = "edgeURL"
	KeyEpsilon            Key = "epsilon"
	KeyEsep               Key = "esep"
	KeyFillColor          Key = "fillcolor"
	KeyFixedSize          Key = "fixedsize"
	KeyFontColor          Key = "fontcolor"
	KeyFontName           Key = "fontname"
	KeyFontNames          Key = "fontnames"
	KeyFontPath           Key = "fontpath"
	KeyFontSize           Key = "fontsize"
	KeyForceLabels        Key = "forcelabels"
	KeyGradientAngle      Key = "gradientangle"
	KeyGroup              Key = "group"
	KeyHeadLp             Key = "head_lp"
	KeyHeadClip           Key = "headclip"
	KeyHeadHref           Key = "headhref"
	KeyHeadLabel          Key = "headlabel"
	KeyHeadPort           Key = "headport"
	KeyHeadTarget         Key = "headtarget"
	KeyHeadTooltip        Key = "headtooltip"
	KeyHeadURL            Key = "headURL"
	KeyHeight             Key = "height"
	KeyHref               Key = "href"
	KeyID                 Key = "id"
	KeyImage              Key = "image"
	KeyImagePath          Key = "imagepath"
	KeyImagePos           Key = "imagepos"
	KeyImageScale         Key = "imagescale"
	KeyInputScale         Key = "inputscale"
	KeyK                  Key = "K"
	KeyLabel              Key = "label"
	KeyLabelScheme        Key = "label_scheme"
	KeyLabelAngle         Key = "labelangle"
	KeyLabelDistance      Key = "labeldistance"
	KeyLabelFloat         Key = "labelfloat"
	KeyLabelFontColor     Key = "labelfontcolor"
	KeyLabelFontName      Key = "labelfontname"
	KeyLabelFontSize      Key = "labelfontsize"
	KeyLabelHref          Key = "labelhref"
	KeyLabelJust          Key = "labeljust"
	KeyLabelLoc           Key = "labelloc"
	KeyLabelTarget        Key = "labeltarget"
	KeyLabelTooltip       Key = "labeltooltip"
	KeyLabelURL           Key = "labelURL"
	KeyLandscape          Key = "landscape"
	KeyLayer              Key = "layer"
	KeyLayerListSep       Key = "layerlistsep"
	KeyLayers             Key = "layers"
	KeyLayerSelect        Key = "layerselect"
	KeyLayerSep           Key = "layersep"
	KeyLayout             Key = "layout"
	KeyLen                Key = "len"
	KeyLevels             Key = "levels"
	KeyLevelsGap          Key = "levelsgap"
	KeyLHead              Key = "lhead"
	KeyLHeight            Key = "lheight"
	KeyLp                 Key = "lp"
	KeyLTail              Key = "ltail"
	KeyLWidth             Key = "lwidth"
	KeyMargin             Key = "margin"
	KeyMaxIter            Key = "maxiter"
	KeyMcLimit            Key = "mclimit"
	KeyMinDist            Key = "mindist"
	KeyMinLen             Key = "minlen"
	KeyMode               Key = "mode"
	KeyModel              Key = "model"
	KeyMosek              Key = "mosek"
	KeyNewRank            Key = "newrank"
	KeyNodeSep            Key = "nodesep"
	KeyNoJustify          Key = "nojustify"
	KeyNormalize          Key = "normalize"
	KeyNoTranslate        Key = "notranslate"
	KeyNsLimit            Key = "nslimit"
	KeyNsLimit1           Key = "nslimit1"
	KeyOrdering           Key = "ordering"
	KeyOrientation        Key = "orientation"
	KeyOutputOrder        Key = "outputorder"
	KeyOverlap            Key = "overlap"
	KeyOverlapScaling     Key = "overlap_scaling"
	KeyOverlapShrink      Key = "overlap_shrink"
	KeyPack               Key = "pack"
	KeyPackMode           Key = "packmode"
	KeyPad                Key = "pad"
	KeyPage               Key = "page"
	KeyPageDir            Key = "pagedir"
	KeyPenColor           Key = "pencolor"
	KeyPenWidth           Key = "penwidth"
	KeyPeripheries        Key = "peripheries"
	KeyPin                Key = "pin"
	KeyPos                Key = "pos"
	KeyQuadTree           Key = "quadtree"
	KeyQuantum            Key = "quantum"
	KeyRank               Key = "rank"
	KeyRankDir            Key = "rankdir"
	KeyRankSep            Key = "ranksep"
	KeyRatio              Key = "ratio"
	KeyRects              Key = "rects"
	KeyRegular            Key = "regular"
	KeyRemincross         Key = "remincross"
	KeyRepulsiveforce     Key = "repulsiveforce"
	KeyResolution         Key = "resolution"
	KeyRoot               Key = "root"
	KeyRotate             Key = "rotate"
	KeyRotation           Key = "rotation"
	KeySameHead           Key = "samehead"
	KeySameTail           Key = "sametail"
	KeySamplePoints       Key = "samplepoints"
	KeyScale              Key = "scale"
	KeySearchSize         Key = "searchsize"
	KeySep                Key = "sep"
	KeyShape              Key = "shape"
	KeyShapeFile          Key = "shapefile"
	KeyShowBoxes          Key = "showboxes"
	KeySides              Key = "sides"
	KeySize               Key = "size"
	KeySkew               Key = "skew"
	KeySmoothing          Key = "smoothing"
	KeySortv              Key = "sortv"
	KeySplines            Key = "splines"
	KeyStart              Key = "start"
	KeyStyle              Key = "style"
	KeyStylesheet         Key = "stylesheet"
	KeyTailLp             Key = "tail_lp"
	KeyTailClip           Key = "tailclip"
	KeyTailHref           Key = "tailhref"
	KeyTailLabel          Key = "taillabel"
	KeyTailPort           Key = "tailport"
	KeyTailTarget         Key = "tailtarget"
	KeyTailTooltip        Key = "tailtooltip"
	KeyTailURL            Key = "tailURL"
	KeyTarget             Key = "target"
	KeyTooltip            Key = "tooltip"
	KeyTrueColor          Key = "truecolor"
	KeyURL                Key = "URL"
	KeyVertices           Key = "vertices"
	KeyViewport           Key = "viewport"
	KeyVoroMargin         Key = "voro_margin"
	KeyWeight             Key = "weight"
	KeyWidth              Key = "width"
	KeyXdotVersion        Key = "xdotversion"
	KeyXlabel             Key = "xlabel"
	KeyXlp                Key = "xlp"
	KeyZ                  Key = "z"
)
