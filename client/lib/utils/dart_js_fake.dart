/// The JavaScript global object, usually `window`.
external JsObject get context;

/// A proxy on a JavaScript object.
///
/// The properties of the JavaScript object are accessible via the `[]` and
/// `[]=` operators. Methods are callable via [callMethod].
class JsObject {
  /// Constructs a JavaScript object from its native [constructor] and returns
  /// a proxy to it.
  external factory JsObject(JsFunction constructor, [List? arguments]);

  /// Constructs a [JsObject] that proxies a native Dart object; _for expert use
  /// only_.
  ///
  /// Use this constructor only if you wish to get access to JavaScript
  /// properties attached to a browser host object, such as a Node or Blob, that
  /// is normally automatically converted into a native Dart object.
  ///
  /// An exception will be thrown if [object] has the type
  /// `bool`, `num`, or `String`.
  external factory JsObject.fromBrowserObject(Object object);

  /// Recursively converts a JSON-like collection of Dart objects to a
  /// collection of JavaScript objects and returns a [JsObject] proxy to it.
  ///
  /// [object] must be a [Map] or [Iterable], the contents of which are also
  /// converted. Maps and Iterables are copied to a new JavaScript object.
  /// Primitives and other transferable values are directly converted to their
  /// JavaScript type, and all other objects are proxied.
  external factory JsObject.jsify(Object object);

  /// Returns the value associated with [property] from the proxied JavaScript
  /// object.
  ///
  /// The type of [property] must be either [String] or [num].
  external dynamic operator [](Object property);

  // Sets the value associated with [property] on the proxied JavaScript
  // object.
  //
  // The type of [property] must be either [String] or [num].
  external void operator []=(Object property, Object? value);
  @override
  int get hashCode => 0;
  @override
  external bool operator ==(Object other);

  /// Returns `true` if the JavaScript object contains the specified property
  /// either directly or though its prototype chain.
  ///
  /// This is the equivalent of the `in` operator in JavaScript.
  external bool hasProperty(Object property);

  /// Removes [property] from the JavaScript object.
  ///
  /// This is the equivalent of the `delete` operator in JavaScript.
  external void deleteProperty(Object property);

  /// Returns `true` if the JavaScript object has [type] in its prototype chain.
  ///
  /// This is the equivalent of the `instanceof` operator in JavaScript.
  external bool instanceof(JsFunction type);
  @override
  /// Returns the result of the JavaScript objects `toString` method.
  external String toString();

  /// Calls [method] on the JavaScript object with the arguments [args] and
  /// returns the result.
  ///
  /// The type of [method] must be either [String] or [num].
  external dynamic callMethod(Object method, [List? args]);
}

/// A proxy on a JavaScript Function object.
class JsFunction extends JsObject {
  /// Returns a [JsFunction] that captures its 'this' binding and calls [f]
  /// with the value of JavaScript `this` passed as the first argument.
  external factory JsFunction.withThis(Function f);

  /// Invokes the JavaScript function with arguments [args]. If [thisArg] is
  /// supplied it is the value of `this` for the invocation.
  external dynamic apply(List args, {thisArg});
}

/// A [List] that proxies a JavaScript array.
class JsArray<E> {
  /// Creates an empty JavaScript array.
  external factory JsArray();

  /// Creates a new JavaScript array and initializes it to the contents of
  /// [other].
  external factory JsArray.from(Iterable<E> other);

  // Methods required by ListMixin

  external E operator [](Object index);

  external void operator []=(Object index, E value);

  external int get length;

  external set length(int length);

  // Methods overridden for better performance

  external void add(E value);

  external void addAll(Iterable<E> iterable);

  external void insert(int index, E element);

  external E removeAt(int index);

  external E removeLast();

  external void removeRange(int start, int end);

  external void setRange(int start, int end, Iterable<E> iterable,
      [int skipCount = 0]);

  external void sort([int Function(E a, E b)? compare]);
}

/// Returns a wrapper around function [f] that can be called from JavaScript
/// using `package:js` JavaScript interop.
///
/// The calling conventions in Dart2Js differ from JavaScript and so, by
/// default, it is not possible to call a Dart function directly. Wrapping with
/// `allowInterop` creates a function that can be called from JavaScript or
/// Dart. The semantics of the wrapped function are still more strict than
/// JavaScript, and the function will throw if called with too many or too few
/// arguments.
///
/// Calling this method repeatedly on a function will return the same result.
external F allowInterop<F extends Function>(F f);

/// Returns a wrapper around function [f] that can be called from JavaScript
/// using `package:js` JavaScript interop, passing JavaScript `this` as the first
/// argument.
///
/// See [allowInterop].
///
/// When called from Dart, `null` will be passed as the first argument.
external Function allowInteropCaptureThis(Function f);
